package factory

import (
	"github.com/ksd-mx/mx-payments/tx-svc/domain/repository"
)

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
