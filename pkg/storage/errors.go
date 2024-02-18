package storage

import "errors"

// Custom error definitions
var (
	ErrKeyExists   = errors.New("key already exists")
	ErrKeyNotFound = errors.New("key not found")
	ErrStorageFail = errors.New("storage operation failed")
)
