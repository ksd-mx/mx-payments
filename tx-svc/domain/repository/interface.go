package repository

type TransactionRepository interface {
	SaveTransaction(
		id string,
		account string,
		amount float64,
		status string,
		errorMessage string) error
}
