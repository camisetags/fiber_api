package database

import (
	transactions "fiber_api/transaction/entities"
	users "fiber_api/user/entities"

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

	db.AutoMigrate(&transactions.Transaction{})
	db.AutoMigrate(&users.User{})

	return db
}
