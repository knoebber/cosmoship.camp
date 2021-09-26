// Package db provides a postgresql connection pool.
package db

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Conn is a global threadsafe database handle.
var Conn *gorm.DB

// Start connects to a postgres database.
func Start(connString string) (err error) {
	Conn, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to start to database: %w", err)
	}

	if err = Conn.AutoMigrate(User{}); err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}

// NotFound returns whether err is wrapping a gorm record not found error.
func NotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
