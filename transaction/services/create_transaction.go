package services

import (
	"errors"
	"fiber_api/transaction/models"
)

// IRepository to receive something like repo interface
type IRepository interface {
	Create(trans *models.Transaction) (*models.Transaction, error)
	GetBalance() models.Balance
}

// CreateTransactionService will handle the domain logic to create transaction
type CreateTransactionService struct{
	Repo IRepository
}

func (c *CreateTransactionService) checksValidBalance(newTransaction models.Transaction) bool {
	balance := c.Repo.GetBalance()
	return balance.Total <= newTransaction.Value
}

// Execute will execute the domain logic of CreateTransactionService
func (c CreateTransactionService) Execute(newTransaction models.Transaction) (*models.Transaction, error) {
	transactType := newTransaction.Type
	if transactType != "income" && transactType != "outcome" {
		return nil, errors.New("Cannot create transaction type different fom income or outcome")
	}

	if transactType == "outcome" && !c.checksValidBalance(newTransaction) {
		return nil, errors.New("Cannot create transaction with invalid balance")
	}

	createdTransaction, err := c.Repo.Create(&newTransaction)

	if err != nil {
		return nil, err
	}

	return createdTransaction, nil
}
