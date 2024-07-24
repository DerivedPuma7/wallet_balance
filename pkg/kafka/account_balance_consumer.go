package kafka

import (
	"fmt"
  "encoding/json"

	"github.com.br/derivedpuma7/balance/internal/usecase/update_balance"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type AccountBalanceUpdateMessageDto struct {
	Name string `json:"Name"`
	Payload struct {
		AccountIDFrom string `json:"AccountIdFrom"`
		AccountIDTo string `json:"AccountIdTo"`
		BalanceAccountIDFrom int `json:"BalanceAccountIdFrom"`
		BalanceAccountIDTo int `json:"BalanceAccountIdTo"`
	} `json:"Payload"`
}

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
  fmt.Println("consuming")
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
    if err != nil {
      panic(err)
    }

    accountBalanceUpdateMessage, err := c.parseMessage(msg.Value)
    if err != nil {
      panic(err)
    }

    updateAccountBalanceFromInput := update_balance.UpdateBalanceInputDto{
      AccountId: accountBalanceUpdateMessage.Payload.AccountIDFrom,
      Balance: float64(accountBalanceUpdateMessage.Payload.BalanceAccountIDFrom),
    }
    updateAccountBalanceToInput := update_balance.UpdateBalanceInputDto{
      AccountId: accountBalanceUpdateMessage.Payload.AccountIDTo,
      Balance:  float64(accountBalanceUpdateMessage.Payload.BalanceAccountIDTo),
    }
    c.UpdateBalanceUseCase.Execute(updateAccountBalanceFromInput)
    c.UpdateBalanceUseCase.Execute(updateAccountBalanceToInput)
    msgChan <- msg
	}
}

func (c *AccountBalanceConsumer) parseMessage(message []byte) (*AccountBalanceUpdateMessageDto, error) {
  var accountBalanceUpdate AccountBalanceUpdateMessageDto

  err := json.Unmarshal(message, &accountBalanceUpdate)
  if err != nil {
    return nil, err
  }

  return &accountBalanceUpdate, nil
}
