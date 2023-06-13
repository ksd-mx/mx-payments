package kafka

import (
	"testing"

	ck "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ksd-mx/mx-payments/adapter/presenter/transaction"
	"github.com/ksd-mx/mx-payments/domain/entity"
	"github.com/ksd-mx/mx-payments/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	expectedOutput := process_transaction.TransactionOutputDTO{
		ID:           "123",
		Status:       entity.REJECTED,
		ErrorMessage: "the maximum amount for a transaction is 1000",
	}

	configMap := ck.ConfigMap{
		"test.mock.num.brokers": 3,
	}

	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKakfaPresenter())
	err := producer.Publish(expectedOutput, []byte("123"), "test")
	assert.Nil(t, err)
}
