package repositories

import (
	"fiber_api/transaction/models"
	
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	// sq "github.com/Masterminds/squirrel"
)

// UserRepoitory repository
type UserRepoitory struct {
	connection *gorm.DB
}


// SetConnection sets connection to be used by repository
func (t UserRepoitory) SetConnection(connection *gorm.DB) UserRepoitory {
	t.connection = connection
	return t
}

func (t UserRepoitory) getConnection() *gorm.DB {
	if t.connection != nil {
		return t.connection
	}
	panic("Connection was not set")
}

// Create will create and returns the created transaction
func (t UserRepoitory) Create(trans *models.Transaction) (*models.Transaction, error) {
	trans.ID = uuid.NewV4()
	result := t.getConnection().Create(trans)

	if result.Error != nil {
		return nil, result.Error
	}

	return trans, nil
}
