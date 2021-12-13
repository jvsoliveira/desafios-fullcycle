package factory

import "github.com/jvsoliveira/desafios-fullcycle/desafio-go/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
