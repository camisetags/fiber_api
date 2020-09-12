package database

import (
	transactionModels "fiber_api/transaction/models"
	userModels "fiber_api/user/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// New returns a new instance of database
func New() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open("fiber_api.db"),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&transactionModels.Transaction{})
	db.AutoMigrate(&userModels.User{})

	return db
}
