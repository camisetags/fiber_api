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

func generatePasswordHash(password string) (string, error) {
	passBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// Execute will execute the domain logic of CreateTransactionService
func (c RegisterUserService) Execute(newUser UserFields, passwordConfirm string) (*models.User, error) {
	if passwordConfirm != newUser.Password {
		return nil, errors.New("Password is not matching")
	}

	newEmail := newUser.Email
	if _, err := c.Repo.FindByEmail(newEmail); err == nil {
		return nil, errors.New("Email already taken")
	}

	hashedPassword, hashErr := generatePasswordHash(newUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	
	dbUser := &models.User{
		Email: newUser.Email,
		Name: newUser.Name,
		Password: hashedPassword,
	}

	c.Repo.Create(dbUser)
	
	return dbUser, nil
}
