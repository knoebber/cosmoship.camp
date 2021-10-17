package models

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/knoebber/cosmoship.camp/db"
	"github.com/knoebber/cosmoship.camp/usererror"
	"gorm.io/gorm"
)

type MemberSource string

const (
	memberSourceAdmin  MemberSource = "admin"
	memberSourceInvite MemberSource = "member invite"

	MemberSourceGuestBook   MemberSource = "guestbook"
	MemberSourceInPerson    MemberSource = "in person"
	MemberSourcePastBooking MemberSource = "past booking"
)

type Member struct {
	ID         int          `gorm:"primaryKey;autoIncrement" json:"id"`
	Email      string       `gorm:"size:255;unique;not null" json:"email" validate:"email"`
	Name       string       `gorm:"size:255" json:"name"`
	PasswordID *int         `gorm:"unique" json:"-"`
	Phone      string       `gorm:"size:64" json:"phone" validate:"omitempty,e164"`
	Approved   bool         `gorm:"type:boolean not null default false" json:"approved"`
	Source     MemberSource `gorm:"size:16" json:"source" validate:"required"`
	CreatedAt  time.Time    `gorm:"type:timestamp;not null" json:"createdAt"`
	UpdatedAt  time.Time    `gorm:"type:timestamp;not null" json:"updatedAt"`

	Password *Password `json:"-"`
}

func (m *Member) passwordID() *int {
	return m.PasswordID
}

func (m *Member) String() string {
	return fmt.Sprintf("member %d", m.ID)
}

func (m *Member) RecordLogin(ip string) error {
	record := &LoginRecord{MemberID: &m.ID, IP: ip}
	if err := db.Conn.Create(record).Error; err != nil {
		return fmt.Errorf("recording login %q for member %d: %w", ip, m.ID, err)
	}
	return nil
}

func (m *Member) BeforeCreate(tx *gorm.DB) error {
	var count int64

	m.Email = strings.ToLower(m.Email)
	if tx.Model(m).Where("email = ?", m.Email).Count(&count); count > 0 {
		return usererror.Format("%q is already registered", m.Email)
	}

	return nil
}

func (m *Member) Get(id int) error {
	m.ID = id
	if err := db.Conn.First(m).Error; err != nil {
		return fmt.Errorf("getting %s: %w", m, err)
	}

	return nil
}

func (m *Member) Create() (interface{}, error) {
	if err := db.Conn.Create(m).Error; err != nil {
		return nil, fmt.Errorf("creating member %q: %w", m.Email, err)
	}

	return m, nil
}

func (m *Member) Delete(id int) error {
	if err := db.Conn.Delete(m, id).Error; err != nil {
		return fmt.Errorf("deleting member %d: %w", id, err)
	}

	return nil
}

func (*Member) Search(values url.Values) (interface{}, error) {
	var result []Member

	if err := db.Conn.Order("created_at desc").Find(&result).Error; err != nil {
		return nil, fmt.Errorf("searching members: %w", err)
	}

	return result, nil
}
