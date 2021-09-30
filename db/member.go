package db

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/knoebber/cosmoship.camp/usererror"
	"gorm.io/gorm"
)

type Member struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string    `gorm:"size:255;unique;not null" json:"email" validate:"email"`
	Name      string    `gorm:"size:255" json:"name"`
	Phone     string    `gorm:"size:64" json:"phone" validate:"omitempty,e164"`
	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null" json:"updatedAt"`
}

func (m *Member) BeforeCreate(tx *gorm.DB) error {
	var count int64

	m.Email = strings.ToLower(m.Email)
	if tx.Model(m).Where("email = ?", m.Email).Count(&count); count > 0 {
		return usererror.New(fmt.Sprintf("%q is already registered", m.Email))
	}

	return nil
}

func (m *Member) Get(id int) error {
	if err := Conn.First(m, id).Error; err != nil {
		return fmt.Errorf("getting user %d: %w", id, err)
	}

	return nil
}

func (m *Member) Create() (interface{}, error) {
	if err := Conn.Create(m).Error; err != nil {
		return nil, fmt.Errorf("creating user %q: %w", m.Email, err)
	}

	return m, nil
}

func (m *Member) Delete(id int) error {
	if err := Conn.Delete(m, id).Error; err != nil {
		return fmt.Errorf("deleting user %d: %w", id, err)
	}

	return nil
}

func (m *Member) Search(values url.Values) (interface{}, error) {
	var result []Member

	if err := Conn.Order("created_at desc").Find(&result).Error; err != nil {
		return nil, fmt.Errorf("searching users:  %w", err)
	}

	return result, nil
}
