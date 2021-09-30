// Package db provides a postgresql connection pool.
package db

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Conn is a global threadsafe database handle.
var Conn *gorm.DB

// Start connects to a postgres database.
func Start(connString string) (err error) {
	Conn, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return fmt.Errorf("failed to start to database: %w", err)
	}

	if err = Conn.AutoMigrate(
		Member{},
	); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}

// NotFound returns whether err is wrapping a gorm record not found error.
func NotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
