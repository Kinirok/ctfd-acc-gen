package gormodel

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	CTFDUser  string `gorm:"not null"`
	CTFDPass  string `gorm:"not null"`
	TeamName  *string
	TeamID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Team struct {
	ID        uint   `gorm:"primaryKey"`
	TeamName  string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
