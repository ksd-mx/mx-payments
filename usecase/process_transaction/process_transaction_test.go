package process_transaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ksd-mx/mx-payments/domain/entity"
	mock_repository "github.com/ksd-mx/mx-payments/domain/repository/mock"
	"github.com/stretchr/testify/assert"
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
		ErrorMessage: "Invalid credit card number",
	}

	controller := gomock.NewController(t)
	defer controller.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(controller)
	repositoryMock.
		EXPECT().
		SaveTransaction(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}
