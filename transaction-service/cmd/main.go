package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ck "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ksd-mx/mx-payments/adapter/broker"
	"github.com/ksd-mx/mx-payments/adapter/broker/kafka"
	"github.com/ksd-mx/mx-payments/adapter/factory"
	"github.com/ksd-mx/mx-payments/adapter/presenter/transaction"
	"github.com/ksd-mx/mx-payments/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	repositoryFactory := factory.NewRepositoryDatabaseFactory(db)
	transactionRepository := repositoryFactory.CreateTransactionRepository()

	configMsgProducer := &ck.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	}
	kafkaPresenter := transaction.NewTransactionKakfaPresenter()
	producer := kafka.NewKafkaProducer(configMsgProducer, kafkaPresenter)

	var msgChan = make(chan *ck.Message)
	configMsgConsumer := &ck.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"client.id":         "mx-payments",
		"group.id":          "mx-payments",
	}
	topics := []string{broker.TransactionTopic}
	consumer := kafka.NewKafkaConsumer(configMsgConsumer, topics)
	go consumer.Consume(msgChan)

	usecase := process_transaction.NewProcessTransaction(transactionRepository, producer, broker.TransactionTopic)

	for msg := range msgChan {
		var input process_transaction.TransactionInputDTO

		json.Unmarshal(msg.Value, &input)
		usecase.Execute(input)
	}
}
