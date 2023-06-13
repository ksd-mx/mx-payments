package entity

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	CreditCard   CreditCard
	Status       string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) SetCreditCard(card CreditCard) {
	t.CreditCard = card
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("the maximum amount for a transaction is 1000")
	}

	if t.Amount < 1 {
		return errors.New("the minimum amount for a transaction is 1")
	}

	return nil
}
