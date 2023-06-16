package kafka

import (
	ck "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ksd-mx/mx-payments/tx-svc/adapter/presenter"
	"github.com/ksd-mx/mx-payments/tx-svc/usecase/process_transaction"
)

type Producer struct {
	ConfigMap *ck.ConfigMap
	Presenter presenter.Presenter[process_transaction.TransactionOutputDTO]
}

func NewKafkaProducer(
	configMap *ck.ConfigMap,
	presenter presenter.Presenter[process_transaction.TransactionOutputDTO],
) *Producer {
	return &Producer{
		ConfigMap: configMap,
		Presenter: presenter,
	}
}

func (p *Producer) Publish(
	message process_transaction.TransactionOutputDTO,
	key []byte,
	toppic string,
) error {
	producer, err := ck.NewProducer(p.ConfigMap)
	if err != nil {
		return err
	}
	err = p.Presenter.Bind(message)
	if err != nil {
		return err
	}
	presenterMessage, err := p.Presenter.Show()
	if err != nil {
		return err
	}
	finalMessage := &ck.Message{
		TopicPartition: ck.TopicPartition{Topic: &toppic, Partition: ck.PartitionAny},
		Value:          presenterMessage,
		Key:            key,
	}
	err = producer.Produce(finalMessage, nil)
	if err != nil {
		return err
	}
	return nil
}
