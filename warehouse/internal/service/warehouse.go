package service

import (
	"context"
	"encoding/json"
	"github.com/Ddarli/app/warehouse/config"
	"github.com/Ddarli/app/warehouse/internal/repo"
	"github.com/Ddarli/app/warehouse/pkg"
	"github.com/Ddarli/utils/kafka"
	"github.com/IBM/sarama"
	"log"
)

type WarehouseService struct {
	producer       kafka.AsyncProducer
	consumer       kafka.ConsumerGroup
	repository     repo.Repository
	messageChannel chan *sarama.ConsumerMessage
}

func New(cfg kafka.ClientConfig) Service {
	client := kafka.NewClient(cfg)

	producer := client.NewAsyncProducer()
	repository := repo.NewWarehouseRepo()

	messageChannel := make(chan *sarama.ConsumerMessage, 100)
	handler := kafka.NewDefaultHandler(messageChannel)
	consumer := client.NewConsumerGroup(config.GroupID, config.Topics, handler)

	return &WarehouseService{
		producer:       producer,
		consumer:       consumer,
		messageChannel: messageChannel,
		repository:     repository,
	}
}

func (s *WarehouseService) GetProducts(ctx context.Context) {
	result, err := s.repository.GetAll(ctx)
	if err != nil {
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return
	}

	message := kafka.Message{
		Topic: "get-products-response",
		Value: jsonData,
	}

	log.Println("Message send: ", message)
	s.producer.SendMessage(ctx, message)
}

func (s *WarehouseService) CheckProducts(ctx context.Context, req *pkg.OrderRequest) {
	result, err := s.repository.CheckQuantity(ctx, req.ID, req.Quantity)
	if err != nil {
		return
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return
	}

	message := kafka.Message{
		Topic: "check-quantity-response",
		Value: jsonData,
	}

	log.Println("Message send: ", message)
	s.producer.SendMessage(ctx, message)

}

func (s *WarehouseService) StartConsuming(ctx context.Context) {
	go func() {
		for {
			select {
			case message, ok := <-s.messageChannel:
				if !ok {
					return
				}
				if message.Topic == "get-products" {
					s.GetProducts(ctx)
				} else if message.Topic == "check-product" {
					var req *pkg.OrderRequest
					err := json.Unmarshal(message.Value, &req)
					if err == nil {
						s.CheckProducts(ctx, req)
					}

				}
			case <-ctx.Done():
				return
			}
		}
	}()

	s.consumer.Consume(ctx)
}
