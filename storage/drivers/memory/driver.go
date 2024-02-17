// memory/MemoryDriver.go
package memory

import (
	"github.com/genovatix/algoliocdm/log"
	"github.com/genovatix/algoliocdm/storage"
	"github.com/genovatix/algoliocdm/storage/drivers"
	"go.uber.org/zap"
	"sync"
)

type MemoryDriver struct {
	mu    sync.RWMutex
	store map[string][]byte
}

func NewMemoryDriver() *MemoryDriver {
	return &MemoryDriver{
		store: make(map[string][]byte),
	}
}

func (m *MemoryDriver) Create(key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.store[key]; exists {
		return drivers.ErrKeyExists
	}
	m.store[key] = value
	// Assuming Logger is a global variable of *zap.Logger
	log.Logger.Info("Create operation successful", zap.String("key", key))
	return nil
}

func (m *MemoryDriver) Read(key string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, exists := m.store[key]
	if !exists {
		log.Logger.Error("Read operation failed", zap.String("key", key), zap.Error(storage.ErrKeyNotFound))
		return nil, storage.ErrKeyNotFound
	}

	log.Logger.Info("Key read", zap.String("key", key))
	return value, nil
}

func (m *MemoryDriver) Update(key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.store[key]; !exists {
		log.Logger.Error("Update operation failed", zap.String("key", key), zap.Error(storage.ErrKeyNotFound))
		return storage.ErrKeyNotFound
	}

	m.store[key] = value
	log.Logger.Info("Key updated", zap.String("key", key))
	return nil
}

func (m *MemoryDriver) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.store[key]; !exists {
		log.Logger.Error("Delete operation failed", zap.String("key", key), zap.Error(storage.ErrKeyNotFound))
		return storage.ErrKeyNotFound
	}

	delete(m.store, key)
	log.Logger.Info("Key deleted", zap.String("key", key))
	return nil
}
