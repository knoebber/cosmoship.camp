package models

import (
	"fmt"
	"net/url"
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

func (a *Admin) passwordID() *int {
	return a.PasswordID
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

func (a *Admin) Create() (interface{}, error) {
	if err := db.Conn.Create(a).Error; err != nil {
		return nil, fmt.Errorf("creating admin %q: %w", a.Email, err)
	}

	return a, nil
}

func (a *Admin) Delete(id int) error {
	if err := db.Conn.Delete(a, id).Error; err != nil {
		return fmt.Errorf("deleting admin %d: %w", id, err)
	}

	return nil
}

func (*Admin) Search(values url.Values) (interface{}, error) {
	var result []Admin

	if err := db.Conn.Order("created_at desc").Find(&result).Error; err != nil {
		return nil, fmt.Errorf("searching admins: %w", err)
	}

	return result, nil
}
