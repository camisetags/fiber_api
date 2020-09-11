package models

import (
	"gorm.io/gorm"
)

// BaseRepository the base repository
type BaseRepository interface {
	getConnection() BaseRepository 
	SetConnection(*gorm.DB) 
}
