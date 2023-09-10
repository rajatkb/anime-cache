package services

import (
	"testing"
)

func TestInsertReadDelete(t *testing.T) {
	// Create a new Store
	store := NewLockedMapStore[int64]()

	// Insert a key-value pair
	key := int64(42)
	value := int64(84)
	store.Put(key, value)

	// Read the key for the value
	result, ok := store.Get(key)
	if !ok || result != value {
		t.Errorf("Expected value %d for key %d, but got %d", value, key, result)
	}

	// Delete the key
	store.Delete(key)

	// Read again and validate it's empty
	result, ok = store.Get(key)
	if ok {
		t.Errorf("Expected key %d to be empty after deletion, but got value %d (ok: %v)", key, result, ok)
	}
}
