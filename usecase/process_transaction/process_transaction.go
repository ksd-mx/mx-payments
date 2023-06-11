package process_transaction

import (
	"github.com/ksd-mx/mx-payments/domain/entity"
	"github.com/ksd-mx/mx-payments/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionInputDTO) (TransactionOutputDTO, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	card, invalidCC := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
	if invalidCC != nil {
		err := p.Repository.SaveTransaction(transaction.ID, transaction.AccountID, transaction.Amount, entity.REJECTED, invalidCC.Error())
		if err != nil {
			return TransactionOutputDTO{}, err
		}

		output := TransactionOutputDTO{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidCC.Error(),
		}

		return output, nil
	}
	transaction.SetCreditCard(*card)
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		err := p.Repository.SaveTransaction(
			transaction.ID,
			transaction.AccountID,
			transaction.Amount,
			entity.REJECTED,
			invalidCC.Error())
		if err != nil {
			return TransactionOutputDTO{}, err
		}

		output := TransactionOutputDTO{
			ID:           transaction.ID,
			Status:       entity.REJECTED,
			ErrorMessage: invalidCC.Error(),
		}

		return output, nil
	}
	return TransactionOutputDTO{}, nil
}
