// Package models provides database models
package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name" validate:"required,min=2"`
	Email     string    `gorm:"size:100;uniqueIndex;not null" json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Posts []Post `gorm:"foreignKey:UserID" json:"posts,omitempty"`
}

// TableName specifies the table name
func (User) TableName() string {
	return "users"
}
