package repositories

import (
	"fiber_api/transaction/models"
	
	"gorm.io/gorm"
	sq "github.com/Masterminds/squirrel"
)

// TransactionRepository repository
type TransactionRepository struct {
	connection *gorm.DB
}

func filterType(transactions []models.Transaction, typee string) []models.Transaction {
	var filtered []models.Transaction
	
	for _, transaction := range transactions {
		if transaction.Type == typee {
			filtered = append(filtered, transaction)
		}
	}

	return filtered
}

func calcTransactionType(transactions []models.Transaction, typee string) uint64 {
	var total uint64
	filtered := filterType(transactions, typee)

	for _, transaction := range filtered {
		total += transaction.Value
	}

	return total
}

// SetConnection sets connection to be used by repository
func (t TransactionRepository) SetConnection(connection *gorm.DB) TransactionRepository {
	t.connection = connection
	return t
}

func (t TransactionRepository) getConnection() *gorm.DB {
	if t.connection != nil {
		return t.connection
	}
	panic("Connection was not set")
}

// All will list all transactions in database
func (t TransactionRepository) All() []models.Transaction {
	transactions := []models.Transaction{}
	t.getConnection().Raw(sq.Select("*").From("transactions").ToSql()).Scan(&transactions)
	
	return transactions
}

// Create will create and returns the created transaction
func (t TransactionRepository) Create(trans *models.Transaction) (*models.Transaction, error) {
	result := t.getConnection().Create(trans)

	if result.Error != nil {
		return nil, result.Error
	}

	return trans, nil
}

// GetBalance will calc the balance
func (t TransactionRepository) GetBalance() models.Balance {
	transactions := []models.Transaction{}
	t.getConnection().Find(&transactions)

	income := calcTransactionType(transactions, "income")
	outcome := calcTransactionType(transactions, "outcome")

	return models.Balance{
		Income: income,
		Outcome: outcome,
		Total: income - outcome,
	}
}
