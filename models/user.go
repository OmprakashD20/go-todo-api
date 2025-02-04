package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName      string    `json:"firstName" gorm:"not null"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email" gorm:"unique;not null"`
	HashedPassword string    `json:"hashedPassword" gorm:"not null"`
	PasswordSalt   string    `json:"passwordSalt" gorm:"not null"`
	Todos          []Todo    `json:"todos" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = time.Now()
	}
	return
}

func (u *User) BeforeUpdate(db *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
