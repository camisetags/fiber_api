package entities

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

// User model
type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
