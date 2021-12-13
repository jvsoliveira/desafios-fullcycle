package process_transaction

import (
	"github.com/jvsoliveira/desafios-fullcycle/desafio-go/domain/entity"
	"github.com/jvsoliveira/desafios-fullcycle/desafio-go/domain/repository"
)

type ProcessTransacion struct {
	Repository repository.TransactionRepository
}

func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransacion {
	return &ProcessTransacion{Repository: repository}
}

func (p *ProcessTransacion) Execute(input TransactionDtoInput) error {

	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount

	err := p.Repository.Insert(
		transaction.ID,
		transaction.AccountID,
		transaction.Amount)

	return err
}
