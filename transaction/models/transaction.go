package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
)

// Transaction model
type Transaction struct {
	gorm.Model
	ID        	uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name  		string 		`json:"name"`
	Type 		string		`json:"type"`
	Value		uint64		`json:"value"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
}

// Balance is the calc of all transaction types
type Balance struct {
	Income uint64	`json:"income"`
	Outcome uint64	`json:"outcome"`
	Total uint64	`json:"total"`
}
