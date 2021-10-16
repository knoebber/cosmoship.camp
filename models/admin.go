package models

import (
	"fmt"
	"time"

	"github.com/knoebber/cosmoship.camp/db"
)

type Admin struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email      string    `gorm:"size:255;unique;not null" json:"email" validate:"email"`
	Name       string    `gorm:"size:255" json:"name"`
	PasswordID *int      `gorm:"unique" json:"-"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null" json:"updatedAt"`

	Password *Password `json:"-"`
}

func (a *Admin) String() string {
	return fmt.Sprintf("admin %d", a.ID)
}

func (a *Admin) RecordLogin(ip string) error {
	record := &LoginRecord{AdminID: &a.ID, IP: ip}
	if err := db.Conn.Create(record).Error; err != nil {
		return fmt.Errorf("recording login %q for admin %d: %w", ip, a.ID, err)
	}
	return nil
}

func (a *Admin) Get(id int) error {
	a.ID = id
	if err := db.Conn.First(a).Error; err != nil {
		return fmt.Errorf("getting %s: %w", a, err)
	}

	return nil
}
