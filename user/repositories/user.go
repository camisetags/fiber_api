package repositories

import (
	"fiber_api/user/entities"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
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
func (t UserRepoitory) Create(user *entities.User) (*entities.User, error) {
	user.ID = uuid.NewV4()
	result := t.getConnection().Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// FindByEmail gets user by email
func (t UserRepoitory) FindByEmail(email string) (*entities.User, error) {
	var user entities.User

	result := t.getConnection().
		Where(&entities.User{Email: email}).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
