package repository

type TransactionRepository interface {
	SaveTransaction(
		id string,
		account string,
		amount string,
		status string,
		errorMessage string) error
}
