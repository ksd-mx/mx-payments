package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Transaction_IsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 100

	assert.Nil(t, transaction.IsValid())
}

func Transaction_IsNotValidAmountGraterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 1001

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "the maximum amount for a transaction is 1000", err.Error())
}

func Transaction_IsNotValidAmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "the minimum amount for a transaction is 1", err.Error())
}
