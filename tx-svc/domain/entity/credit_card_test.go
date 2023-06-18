package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCrediCardNumber(t *testing.T) {
	_, err := NewCreditCard("0000000000000000", "John Doe", 12, time.Now().Year(), "123")
	assert.Equal(t, "Invalid credit card number", err.Error())

	_, err = NewCreditCard("4153281702402135", "John Doe", 12, time.Now().Year(), "123")
	assert.Nil(t, err)
}

func TestCrediCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4153281702402135", "John Doe", 0, time.Now().Year(), "123")
	assert.Equal(t, "Invalid expiration month", err.Error())

	_, err = NewCreditCard("4153281702402135", "John Doe", 13, time.Now().Year(), "123")
	assert.Equal(t, "Invalid expiration month", err.Error())

	_, err = NewCreditCard("4153281702402135", "John Doe", 12, time.Now().Year(), "123")
	assert.Nil(t, err)
}

func TestCrediCardExpirationYear(t *testing.T) {
	lastYear := time.Now().AddDate(-1, 0, 0).Year()
	_, err := NewCreditCard("4153281702402135", "John Doe", 12, lastYear, "123")
	assert.Equal(t, "Invalid expiration year", err.Error())

	_, err = NewCreditCard("4153281702402135", "John Doe", 12, time.Now().Year(), "123")
	assert.Nil(t, err)
}

func TestCrediCardCvv(t *testing.T) {
	_, err := NewCreditCard("4153281702402135", "John Doe", 12, time.Now().Year(), "23")
	assert.Equal(t, "Invalid CVV", err.Error())

	_, err = NewCreditCard("4153281702402135", "John Doe", 12, time.Now().Year(), "123")
	assert.Nil(t, err)
}
