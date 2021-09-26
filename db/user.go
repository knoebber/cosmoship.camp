package db

import (
	"fmt"
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) Get(id int) error {
	if err := Conn.First(u, id).Error; err != nil {
		return fmt.Errorf("getting user %d: %w", id, err)
	}

	return nil
}

func (u *User) Create() (interface{}, error) {
	if err := Conn.Create(u).Error; err != nil {
		return nil, fmt.Errorf("creating user %q: %w", u.Email, err)
	}

	return u, nil
}
