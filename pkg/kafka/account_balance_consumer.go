package kafka

import (
	"fmt"

	"github.com.br/derivedpuma7/balance/internal/usecase/update_balance"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type AccountBalanceConsumer struct {
  ConfigMap *ckafka.ConfigMap
	Topics []string
  UpdateBalanceUseCase *update_balance.UpdateBalanceUseCase
}

func NewAccountBalanceConsumer(configMap *ckafka.ConfigMap, topics []string, updateBalanceUseCase *update_balance.UpdateBalanceUseCase) *AccountBalanceConsumer {
	return &AccountBalanceConsumer{
		ConfigMap: configMap,
		Topics:    topics,
    UpdateBalanceUseCase: updateBalanceUseCase,
	}
}

func (c *AccountBalanceConsumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		panic(err)
	}
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
    fmt.Println(msg)
		if err == nil {
			msgChan <- msg
		}
	}
}
