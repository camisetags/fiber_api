package services

import (
	"errors"
	"fiber_api/transaction/models"
)

// IRepository to receive something like repo interface
type IRepository interface {
	Create(trans models.Transaction) (*models.Transaction, error)
}

// CreateTransactionService will handle the domain logic to create transaction
type CreateTransactionService struct{
	Repo IRepository
}

// CreateTransactionDTO params to create transaction service
type CreateTransactionDTO struct {
	Transaction models.Transaction
}

// Execute will execute the domain logic of CreateTransactionService
func (l CreateTransactionService) Execute(param CreateTransactionDTO) (*models.Transaction, error) {
	transactType := param.Transaction.Type
	if transactType != "income" && transactType != "outcome" {
		return nil, errors.New("Cannot create transaction type different fom income or outcome")
	}

	createdTransaction, err := l.Repo.Create(param.Transaction)

	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}
