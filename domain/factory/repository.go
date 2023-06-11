package factory

import "github.com/ksd-mx/mx-payments/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
