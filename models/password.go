package models

import (
	"fmt"
	"time"

	"github.com/knoebber/cosmoship.camp/db"
	"github.com/knoebber/cosmoship.camp/usererror"
	"golang.org/x/crypto/bcrypt"
)

const minPasswordLength = 8

type Password struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Hash      []byte    `gorm:"type:char(60)" json:"-"`
	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"createdAt"`
}

func (p *Password) String() string {
	return fmt.Sprintf("password %d", p.ID)
}

func hashPassword(password string) ([]byte, error) {
	if len(password) < minPasswordLength {
		return nil, usererror.Format("password must be at least %d characters", minPasswordLength)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}

	return passwordHash, nil
}

// Login logs either a member or an admin in.
func checkPassword(email, password string) (User, error) {
	var err error

	admin := new(Admin)

	// First look for an admin with email.
	if err = db.Conn.Joins("Password").First(admin, "email = ?", email).Error; err == nil {
		return admin, bcrypt.CompareHashAndPassword(admin.Password.Hash, []byte(password))
	} else if db.NotFound(err) {
		// When admin with email isn't found, look for member.
		member := new(Member)
		if err = db.Conn.Joins("Password").First(member, "email = ?", email).Error; err == nil {
			return member, bcrypt.CompareHashAndPassword(member.Password.Hash, []byte(password))
		}
	}

	return nil, err
}
