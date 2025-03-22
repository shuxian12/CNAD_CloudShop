package domain

import (
	"strings"
)

type User struct {
	Username string
}

func NewUser(username string) *User {
	return &User{Username: strings.ToLower(username)}
}