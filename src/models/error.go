package models

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrExpired       = errors.New("short url expired")
)
