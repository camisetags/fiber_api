package database

import (
	"fiber_api/transaction/models"

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

	db.AutoMigrate(&models.Transaction{})

	return db
}
