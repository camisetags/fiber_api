package repositories

import (
	"fiber_api/transaction/entities"

	sq "github.com/Masterminds/squirrel"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// TransactionRepository repository
type TransactionRepository struct {
	connection *gorm.DB
}

func filterType(transactions []entities.Transaction, typee string) []entities.Transaction {
	var filtered []entities.Transaction

	for _, transaction := range transactions {
		if transaction.Type == typee {
			filtered = append(filtered, transaction)
		}
	}

	return filtered
}

func calcTransactionType(transactions []entities.Transaction, typee string) uint64 {
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
func (t TransactionRepository) All() []entities.Transaction {
	transactions := []entities.Transaction{}
	t.getConnection().Raw(sq.Select("*").From("transactions").ToSql()).Scan(&transactions)

	return transactions
}

// Create will create and returns the created transaction
func (t TransactionRepository) Create(trans *entities.Transaction) (*entities.Transaction, error) {
	trans.ID = uuid.NewV4()
	result := t.getConnection().Create(trans)

	if result.Error != nil {
		return nil, result.Error
	}

	return trans, nil
}

// GetBalance will calc the balance
func (t TransactionRepository) GetBalance() entities.Balance {
	transactions := []entities.Transaction{}
	t.getConnection().Find(&transactions)

	income := calcTransactionType(transactions, "income")
	outcome := calcTransactionType(transactions, "outcome")

	return entities.Balance{
		Income:  income,
		Outcome: outcome,
		Total:   income - outcome,
	}
}
