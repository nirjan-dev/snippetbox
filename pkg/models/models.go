package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matched record found")
var ErrDuplicateEmail = errors.New("models: duplicate email")
var ErrInvalidCredentials = errors.New("models: invalid credentials")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
