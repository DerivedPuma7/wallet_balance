package main

import (
	"database/sql"
	"fmt"

	"github.com.br/derivedpuma7/balance/internal/database"
	"github.com.br/derivedpuma7/balance/internal/usecase/get_balance_by_account"
	"github.com.br/derivedpuma7/balance/internal/usecase/update_balance"
	"github.com.br/derivedpuma7/balance/internal/web"
	"github.com.br/derivedpuma7/balance/internal/web/webserver"
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
  getBalaceUseCase := get_balance_by_account.NewGetBalanceByAccountUseCase(accountBalanceDb)
  
  httpPort := ":3003"
  webserver := webserver.NewWebServer(httpPort)
  balanceHandler := web.NewWebBalanceHandler(*getBalaceUseCase)
  webserver.AddHandler("/balances/{account_id}", balanceHandler.Handle)
  fmt.Println("server is running on: http://localhost", httpPort)
  go webserver.Start()

  configMap := ckafka.ConfigMap{
    "bootstrap.servers": "kafka:29092",
    "group.id": "wallet",
    "auto.offset.reset": "earliest",
  }
  kafkaTopics := []string{"balances"}

  kafkaConsumer := kafka.NewAccountBalanceConsumer(&configMap, kafkaTopics, updateBalanceUseCase)
  channel := make(chan *ckafka.Message)
  go kafkaConsumer.Consume(channel)
  for msg := range channel{
    fmt.Println(string(msg.Value))
  }

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
