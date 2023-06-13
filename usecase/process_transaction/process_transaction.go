package process_transaction

import (
	"github.com/ksd-mx/mx-payments/adapter/broker"
	"github.com/ksd-mx/mx-payments/domain/entity"
	"github.com/ksd-mx/mx-payments/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
	Producer   broker.ProducerInterface[TransactionOutputDTO]
	Topic      string
}

func NewProcessTransaction(
	repository repository.TransactionRepository,
	producer broker.ProducerInterface[TransactionOutputDTO],
	topic string,
) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository, Producer: producer, Topic: topic}
}

func (p *ProcessTransaction) Execute(input TransactionInputDTO) (TransactionOutputDTO, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	card, invalidCC := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
	if invalidCC != nil {
		return p.rejectTransaction(transaction, invalidCC)
	}
	transaction.SetCreditCard(*card)
	invalidTransaction := transaction.IsValid()
	if invalidTransaction != nil {
		return p.rejectTransaction(transaction, invalidTransaction)
	}
	return p.approveTransaction(transaction)
}

func (p *ProcessTransaction) approveTransaction(transaction *entity.Transaction) (TransactionOutputDTO, error) {
	err := p.Repository.SaveTransaction(
		transaction.ID,
		transaction.AccountID,
		transaction.Amount,
		entity.APPROVED,
		"",
	)
	if err != nil {
		return TransactionOutputDTO{}, err
	}

	output := TransactionOutputDTO{
		ID:           transaction.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	err = p.publish(output, []byte(transaction.ID))

	if err != nil {
		return TransactionOutputDTO{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionOutputDTO, error) {
	err := p.Repository.SaveTransaction(
		transaction.ID,
		transaction.AccountID,
		transaction.Amount,
		entity.REJECTED,
		invalidTransaction.Error(),
	)
	if err != nil {
		return TransactionOutputDTO{}, err
	}

	output := TransactionOutputDTO{
		ID:           transaction.ID,
		Status:       entity.REJECTED,
		ErrorMessage: invalidTransaction.Error(),
	}

	err = p.publish(output, []byte(transaction.ID))

	if err != nil {
		return TransactionOutputDTO{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) publish(output TransactionOutputDTO, key []byte) error {
	err := p.Producer.Publish(output, []byte(output.ID), p.Topic)
	if err != nil {
		return err
	}
	return nil
}
