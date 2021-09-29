package db

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/knoebber/cosmoship.camp/usererror"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"size:255;unique;not null" json:"email" validate:"email"`
	Name      string    `gorm:"size:255" json:"name"`
	Phone     string    `gorm:"size:64" json:"phone" validate:"omitempty,e164"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var count int64

	u.Email = strings.ToLower(u.Email)
	if tx.Model(u).Where("email = ?", u.Email).Count(&count); count > 0 {
		return usererror.New(fmt.Sprintf("%q is already registered", u.Email))
	}

	return nil
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

func (u *User) Search(values url.Values) (interface{}, error) {
	var result []User

	if err := Conn.Find(&result).Error; err != nil {
		return nil, fmt.Errorf("searching users:  %w", err)
	}

	return result, nil
}
