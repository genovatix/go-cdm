package drivers

import "errors"

// StorageDriver defines the interface for a storage backend.
type StorageDriver interface {
	Create(key string, value []byte) error
	Read(key string) ([]byte, error)
	Update(key string, value []byte) error
	Delete(key string) error
}

// Common storage errors.
var (
	ErrKeyExists   = errors.New("key already exists")
	ErrKeyNotFound = errors.New("key not found")
)
