package transaction

import (
	"encoding/json"

	"github.com/ksd-mx/mx-payments/usecase/process_transaction"
)

type KafkaPresenter struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewTransactionKakfaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (p *KafkaPresenter) Bind(transaction process_transaction.TransactionOutputDTO) error {
	p.ID = transaction.ID
	p.Status = transaction.Status
	p.ErrorMessage = transaction.ErrorMessage
	return nil
}

func (p *KafkaPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return j, nil
}
