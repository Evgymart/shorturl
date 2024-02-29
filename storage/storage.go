package storage

import "errors"

var (
	ErrUrlNotFound      = errors.New("url is not found")
	ErrUrlAlreadyExists = errors.New("url already exists")
)
