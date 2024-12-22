package config

import (
	"github.com/Ddarli/utils/kafka"
	"time"
)

const (
	GroupID       = "SOso"
	Key           = "secret-key"
	TokenLifeTime = time.Duration(1) * time.Hour
)

var Topics = []string{"get-products-response", "check-quantity-response", "make-order-response"}

var Cfg = kafka.ClientConfig{
	Brokers:            []string{"kafka:9092"},
	InsecureSkipVerify: true,
	Producer: &kafka.ProducerConfig{
		RequireAcks: 1,
		MaxAttempts: 1,
		Compression: 1,
		RetryMax:    4,
		Idempotent: struct {
			Mode            bool
			MaxOpenRequests int
			RetryMax        int
		}{false, 1, 1},
	},
	Consumer: &kafka.ConsumerConfig{
		Addresses:    []string{"kafka:9092"},
		Assignor:     "round-robin",
		OffsetNewest: false,
		AutoCommit:   true,
	},
}
