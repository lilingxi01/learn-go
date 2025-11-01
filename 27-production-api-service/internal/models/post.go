// Package models provides database models
package models

import (
	"time"
)

// Post represents a blog post
type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"size:200;not null" json:"title" validate:"required"`
	Content   string    `gorm:"type:text" json:"content"`
	Published bool      `gorm:"default:false" json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// TableName specifies the table name
func (Post) TableName() string {
	return "posts"
}
