package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/knoebber/cosmoship.camp/redispool"
)

type User interface {
	Getter
	String() string
	RecordLogin(string) error
	passwordID() *int
}

func CheckSession(sessionKey string) (sessionValue string, err error) {
	rc := redispool.Get()
	defer rc.Close()

	sessionValue, err = redis.String(rc.Do("GETEX", sessionKey, "EX", sessionDurationSeconds))
	if err != nil {
		return "", fmt.Errorf("getting session %q: %w", sessionKey, err)
	}

	return
}

// ParseSessionValue finds the user's ID and whether they are an admin.
// See .String() methods for types that implement the User interface for more information.
func ParseSessionValue(sessionValue string) (id int, isAdmin bool) {
	parts := strings.Split(sessionValue, " ")
	isAdmin = parts[0] == "admin"
	id, _ = strconv.Atoi(parts[1])
	return
}

func Login(email, password, ip string) (sessionKey string, err error) {
	var user User

	user, err = checkPassword(email, password)
	if err != nil {
		return "", err
	}
	if err = user.RecordLogin(ip); err != nil {
		return "", err
	}

	return setSession(user.String())
}

func setSession(value string) (sessionKey string, err error) {
	sessionKey, err = makeSessionKey()
	if err != nil {
		return "", err
	}

	rc := redispool.Get()
	defer rc.Close()
	if _, err := rc.Do("SET", sessionKey, value, "EX", sessionDurationSeconds); err != nil {
		return "", fmt.Errorf("setting session %q: %w", value, err)
	}

	return
}

func makeSessionKey() (string, error) {
	buff, err := randomBytes(sessionLength)
	if err != nil {
		return "", fmt.Errorf("generating session: %w", err)
	}

	return fmt.Sprintf("session:%s", base64.URLEncoding.EncodeToString(buff)), nil
}

func randomBytes(n int) ([]byte, error) {
	buff := make([]byte, n)

	if _, err := io.ReadFull(rand.Reader, buff); err != nil {
		return nil, err
	}

	return buff, nil
}
