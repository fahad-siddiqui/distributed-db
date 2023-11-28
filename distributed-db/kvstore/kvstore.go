package kvstore

import (
    "sync"
)

// KVStore represents the distributed key-value store.
type KVStore struct {
    mu      sync.RWMutex
    data    map[string]string
}

// NewKVStore initializes a new KVStore.
func NewKVStore() *KVStore {
    return &KVStore{
        data: make(map[string]string),
    }
}

// Retrieves the value for a given key.
func (kv *KVStore) Get(key string) (string, bool) {
    kv.mu.RLock()
    defer kv.mu.RUnlock()
    value, exists := kv.data[key]
    return value, exists
}


//Get all values present in key value store
func (kv *KVStore) GetAllValues() []string {
    kv.mu.RLock()
    defer kv.mu.RUnlock()

    values := make([]string, 0, len(kv.data))
    for _, value := range kv.data {
        values = append(values, value)
    }

    return values
}

// GetAllKeyValues returns all key-value pairs in the key-value store.
func (kv *KVStore) GetAllKeyValues() map[string]string {
    kv.mu.RLock()
    defer kv.mu.RUnlock()

    keyValues := make(map[string]string)
    for key, value := range kv.data {
        keyValues[key] = value
    }
    return keyValues
}

// Set updates the value for a given key.
func (kv *KVStore) Set(key, value string) {
    kv.mu.Lock()
    defer kv.mu.Unlock()
    kv.data[key] = value
}

func (kv *KVStore) Delete(key string) {
    kv.mu.Lock()
    defer kv.mu.Unlock()
    delete(kv.data, key)
}