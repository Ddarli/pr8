package config

import "github.com/Ddarli/utils/kafka"

const (
	GroupID  = "SOso"
	Host     = "postgres"
	Port     = 5432
	User     = "myuser"
	Password = "mypassword"
	Dbname   = "mydb"
)

var Topics = []string{"make-order"}

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
