package services

import (
	"errors"
	"fiber_api/user/models"
	
	"golang.org/x/crypto/bcrypt"
)

// IRepository to receive something like repo interface
type IRepository interface {
	Create(*models.User) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

// UserFields basic struct of user to create a new one
type UserFields struct {
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
	NewUser				 UserFields
}

func generatePasswordHash(password string) (string, error) {
	passBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// Execute will execute the domain logic of CreateTransactionService
func (c RegisterUserService) Execute(params UserRegisterDTO) (*models.User, error) {
	if params.PasswordConfirmation != params.NewUser.Password {
		return nil, errors.New("Password is not matching")
	}

	newEmail := params.NewUser.Email
	if _, err := c.Repo.FindByEmail(newEmail); err == nil {
		return nil, errors.New("Email already taken")
	}

	hashedPassword, hashErr := generatePasswordHash(params.NewUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	
	newUser := &models.User{
		Email: params.NewUser.Email,
		Name: params.NewUser.Name,
		Password: hashedPassword,
	}

	c.Repo.Create(newUser)
	
	return newUser, nil
}
