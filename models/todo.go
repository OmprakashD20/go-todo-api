package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Priority    string    `json:"priority" gorm:"not null"`
	DueDate     time.Time `json:"dueDate" gorm:"type:timestamptz;not null"`
	IsCompleted bool      `json:"isCompleted" gorm:"default:false"`
	UserID      uint      `json:"userId" gorm:"not null"`
	// User        User      `json:"-"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (u *Todo) BeforeCreate(db *gorm.DB) (err error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = time.Now()
	}
	return
}

func (u *Todo) BeforeUpdate(db *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
