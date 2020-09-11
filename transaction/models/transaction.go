package models

import (
	"time"
	"gorm.io/gorm"
)

// Transaction model
type Transaction struct {
	gorm.Model
	ID			uint		`gorm:"primaryKey"`
	Name  		string 
	Type 		string
	Value		uint64
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

// Balance is the calc of all transaction types
type Balance struct {
	Income uint64
	Outcome uint64
	Total uint64
}
