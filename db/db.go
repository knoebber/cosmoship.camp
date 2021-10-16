package db

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Conn is a global threadsafe database handle.
var Conn *gorm.DB

// Start connects to a postgres database.
func Start(connString string) (err error) {
	Conn, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger:  logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return fmt.Errorf("failed to start to database: %w", err)
	}

	return nil
}

// NotFound returns whether err is wrapping a gorm record not found error.
func NotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

// QueryError returns whether the error is an unexpected sql error, as opposed to a record not found.
func QueryError(err error) bool {
	return err != nil && !errors.Is(err, gorm.ErrRecordNotFound)
}
