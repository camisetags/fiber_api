package services

import (
	"errors"
	"fiber_api/user/models"
)

// IRepository to receive something like repo interface
type IRepository interface {
	Create(trans *models.User) (*models.User, error)
}

// UserCreation basic struct of user to create a new one
type UserCreation struct {
	Name string
	Email string
	Password string
}

// RegisterUserService will handle the domain logic to create transaction
type RegisterUserService struct{
	Repo IRepository
}

// UserRegisterDTO data access to user register params
type UserRegisterDTO struct {
	PasswordConfirmation string
	NewUser				 UserCreation
}

// Execute will execute the domain logic of CreateTransactionService
func (c RegisterUserService) Execute(params UserRegisterDTO) (*models.User, error) {
	if params.PasswordConfirmation != params.NewUser.Password {
		return nil, errors.New("Password is not matching")
	}

	return &params.NewUser, nil
}
