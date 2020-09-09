package repositories

import (
	"fiber_api/transaction/models"
	
	"gorm.io/gorm"
)

// TransactionRepository repository
type TransactionRepository struct {
	Connection *gorm.DB
}

// All will list all transactions in database
func (t TransactionRepository) All() []models.Transaction {
	transactions := []models.Transaction{}
	t.Connection.Find(&transactions)
	
	return transactions
}

// Create will create and returns the created transaction
func (t TransactionRepository) Create(trans models.Transaction) (*models.Transaction, error) {
	result := t.Connection.Create(&trans)

	if result.Error != nil {
		return nil, result.Error
	}

	return &trans, nil
}
