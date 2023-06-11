package process_transaction

import (
	"testing"
	"time"

	"github.com/ksd-mx/mx-payments/domain/entity"
)

func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionInputDTO{
		ID:                        "1",
		AccountID:                 "1",
		CreditCardNumber:          "1234567890123456",
		CreditCardName:            "John Doe",
		CreditCardExpirationMonth: 01,
		CreditCardExpirationYear:  time.Now().Year(),
		CreditCardCVV:             "123",
		Amount:                    100,
	}

	expectedOutput := TransactionOutputDTO{
		ID:           "1",
		Status:       entity.REJECTED,
		ErrorMessage: "invalid credit card",
	}

	usecase := NewProcessTransaction(repository)
}
