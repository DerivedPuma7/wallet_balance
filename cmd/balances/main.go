package main

import (
	"database/sql"
	"fmt"

	"github.com.br/derivedpuma7/balance/internal/database"
	"github.com.br/derivedpuma7/balance/internal/usecase/update_balance"
	"github.com.br/derivedpuma7/balance/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql-balances", "3306", "balance"))
  if err != nil {
    panic(err)
  }
  defer db.Close()
  startDatabase(db)

  accountBalanceDb := database.NewAccountBalanceDb(db)
  updateBalanceUseCase := update_balance.NewUpdateBalanceUseCase(accountBalanceDb)

  configMap := ckafka.ConfigMap{
    "bootstrap.servers": "kafka:29092",
    "group.id": "wallet",
  }
  kafkaTopics := []string{"balances"}

  kafkaConsumer := kafka.NewAccountBalanceConsumer(&configMap, kafkaTopics, updateBalanceUseCase)
  channel := make(chan *ckafka.Message)
  kafkaConsumer.Consume(channel)
}

func startDatabase(db *sql.DB) {
  sql := `
    CREATE TABLE IF NOT EXISTS balances (
      id varchar(255),
      accountId varchar(255),
      balance float,
      updatedAt date
    );`

  _, err := db.Exec(sql)
  if err != nil {
    panic(err)
  }
}
