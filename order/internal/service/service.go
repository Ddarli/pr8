package service

import (
	"context"
	"encoding/json"
	"github.com/Ddarli/app/order/config"
	"github.com/Ddarli/app/order/internal/repo"
	"github.com/Ddarli/app/order/pkg/models"
	"github.com/Ddarli/utils/kafka"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type OrderService struct {
	producer       kafka.AsyncProducer
	consumer       kafka.ConsumerGroup
	messageChannel chan *sarama.ConsumerMessage
	repository     repo.Repository
}

func New(cfg kafka.ClientConfig) Service {
	client := kafka.NewClient(cfg)

	producer := client.NewAsyncProducer()

	messageChannel := make(chan *sarama.ConsumerMessage, 100)
	handler := kafka.NewDefaultHandler(messageChannel)
	consumer := client.NewConsumerGroup(config.GroupID, config.Topics, handler)
	rep := repo.New()

	return &OrderService{
		producer:       producer,
		consumer:       consumer,
		messageChannel: messageChannel,
		repository:     rep,
	}
}

func (s *OrderService) StartConsuming(ctx context.Context) {
	go func() {
		for {
			select {
			case message, ok := <-s.messageChannel:
				if !ok {
					return
				}
				request := models.OrderRequest{}

				err := json.Unmarshal(message.Value, &request)
				if err != nil {
					log.Println(err)
				}
				response, err := s.processOrder(ctx, &request)
				if err != nil {
					// log error
					continue
				}
				responseBytes, err := json.Marshal(response)
				if err != nil {
					// log error
					continue
				}
				msg := kafka.Message{
					Topic: "make-order-response",
					Value: responseBytes,
				}
				s.producer.SendMessage(ctx, msg)
			case <-ctx.Done():
				return
			}
		}
	}()

	s.consumer.Consume(ctx)
}

func (s *OrderService) processOrder(ctx context.Context, request *models.OrderRequest) (*models.OrderResponse, error) {
	response := &models.OrderResponse{}
	order := &models.Order{
		Customer: "Customer",
		Date:     time.Now(),
		Total:    float32(request.Quantity),
	}

	res, err := s.repository.SaveOrder(ctx, order)
	response.Success = res

	return response, err
}
