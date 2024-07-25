## Balance microservice - Event Driven Architecture

## Responsability:
- It consumes an Apache Kafka topic and updates balance for the accounts involved in a transaction

- It exposes an HTTP endpoint responsable for fetching an account balance

## Seting up environment
- It uses a specific docker network to comunicate with other containers, so before "compose up", make sure to create the network correctly
  > docker network create wallet-network

- Once the network is created, run containers
  > docker compose up -d

- Make sure Apache Kafka topics were created
  - go to http://localhost:9021 for Control Center or http://localhost:8000 for Kafka Ui  
    > create both 'balances' and 'transactions' topics, with default configs

## Start application
- access go container and run main.go
  > docker exec -it wallet-balances bash
  > go run cmd/balances/main.go
