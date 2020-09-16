package entities

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

// Transaction model
type Transaction struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Value     uint64    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
