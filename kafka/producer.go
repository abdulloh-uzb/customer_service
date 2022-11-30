package messagebroker

import (
	"context"
	"encoding/json"
	"exam/customer_service/config"
	pbc "exam/customer_service/genproto/customer"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Conn      *kafka.Conn
	CloseConn func()
}

func NewProducer(cfg config.Config) (*Producer, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", cfg.KafkaHost+":"+cfg.KafkaPort, cfg.KafkaTopic, 0)
	if err != nil {
		fmt.Println("Error is here")
		return &Producer{}, err
	}
	return &Producer{
		Conn: conn,
		CloseConn: func() {
			conn.Close()
		},
	}, nil
}

func (p *Producer) ProducerCreate(message *pbc.CustomerRequest) error {
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = p.Conn.WriteMessages(kafka.Message{
		Value: value,
	})
	if err != nil {
		fmt.Println("====>", err)
		return err
	}
	return nil
}
