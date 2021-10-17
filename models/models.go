package models

import (
	"fmt"
	"net/url"

	"github.com/knoebber/cosmoship.camp/db"
)

const (
	sessionLength = 24
	// sessionDurationSeconds = 60 * 60 * 24 * 7
	sessionDurationSeconds = 60 * 60
)

type Getter interface {
	Get(id int) error
}

type Creater interface {
	Create() (interface{}, error)
}

type Deleter interface {
	Delete(id int) error
}

type Searcher interface {
	Search(url.Values) (interface{}, error)
}

func Migrate() error {
	if err := db.Conn.AutoMigrate(
		new(Admin),
		new(Member),
		new(Password),
		new(LoginRecord),
		new(InvitationReference),
	); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}
	return nil
}
