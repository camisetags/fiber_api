package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
)

// User model
type User struct {
	gorm.Model
	ID        	uuid.UUID 	`gorm:"type:uuid;primary_key;"`
	Name  		string 		`json:"name"`
	Email  		string 		`json:"email"`
	password	string		
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
}
