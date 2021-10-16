package redispool

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

const timeout = 15 * time.Second

var pool *redis.Pool

func Start(connString string) error {
	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: timeout,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", connString) },
	}

	conn := pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return fmt.Errorf("failed to ping redis: %w", err)
	}

	return nil
}

func Get() redis.Conn {
	return pool.Get()
}
