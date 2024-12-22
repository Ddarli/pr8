package service

import (
	"context"
	"encoding/json"
	"github.com/Ddarli/app/shop/config"
	pkg "github.com/Ddarli/app/shop/pkg/models"
	"github.com/Ddarli/utils/kafka"
	"github.com/Ddarli/utils/models"
	"github.com/IBM/sarama"
)

type ShopService struct {
	producer       kafka.AsyncProducer
	consumer       kafka.ConsumerGroup
	messageChannel chan *sarama.ConsumerMessage
	pr             chan []*models.Product
	qResp          chan bool
	orders         chan *pkg.OrderResponse
}

func New(cfg kafka.ClientConfig) Service {
	client := kafka.NewClient(cfg)

	producer := client.NewAsyncProducer()

	messageChannel := make(chan *sarama.ConsumerMessage, 100)
	handler := kafka.NewDefaultHandler(messageChannel)
	consumer := client.NewConsumerGroup(config.GroupID, config.Topics, handler)
	qResp := make(chan bool)
	pr := make(chan []*models.Product)
	orders := make(chan *pkg.OrderResponse)

	return &ShopService{
		producer:       producer,
		consumer:       consumer,
		messageChannel: messageChannel,
		pr:             pr,
		qResp:          qResp,
		orders:         orders,
	}
}

func (s *ShopService) GetAll(ctx context.Context) ([]*models.Product, error) {
	products := s.productsRequest(ctx)

	return products, nil
}

func (s *ShopService) StartConsuming(ctx context.Context) {
	go func() {
		for {
			select {
			case message, ok := <-s.messageChannel:
				if !ok {
					return
				}
				if message.Topic == "get-products-response" {
					var prods []*models.Product
					err := json.Unmarshal(message.Value, &prods)
					if err == nil {
						s.pr <- prods
					}
				} else if message.Topic == "check-quantity-response" {
					var resp bool
					err := json.Unmarshal(message.Value, &resp)
					if err == nil {
						s.qResp <- resp
					}
				} else if message.Topic == "make-order-response" {
					var resp *pkg.OrderResponse
					err := json.Unmarshal(message.Value, &resp)
					if err == nil {
						s.orders <- resp
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	s.consumer.Consume(ctx)
}

func (s *ShopService) ProcessOrder(ctx context.Context, request pkg.OrderRequest) (*pkg.OrderResponse, error) {
	var response *pkg.OrderResponse

	available, err := s.checkQuantity(ctx, &request)
	if available && err == nil {
		response, err = s.makeOrder(ctx, &request)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return response, err
}

func (s *ShopService) productsRequest(ctx context.Context) []*models.Product {
	message := kafka.Message{
		Topic: "get-products",
	}
	s.producer.SendMessage(ctx, message)

	select {
	case products := <-s.pr:
		return products
	case <-ctx.Done():
		return nil
	}
}

func (s *ShopService) checkQuantity(ctx context.Context, request *pkg.OrderRequest) (bool, error) {
	val, err := json.Marshal(request)
	if err != nil {
		return false, err
	}

	message := kafka.Message{
		Topic: "check-product",
		Value: val,
	}
	s.producer.SendMessage(ctx, message)

	select {
	case response := <-s.qResp:
		return response, nil
	case <-ctx.Done():
		return false, nil
	}
}

func (s *ShopService) makeOrder(ctx context.Context, request *pkg.OrderRequest) (*pkg.OrderResponse, error) {
	val, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	message := kafka.Message{
		Topic: "make-order",
		Value: val,
	}
	s.producer.SendMessage(ctx, message)

	select {
	case response := <-s.orders:
		return response, nil
	case <-ctx.Done():
		return nil, nil
	}
}
