package services

import (
	"animecache/entities"
	"sync"
)

type LockedMapStore[V any] struct {
	hashMap map[int64]V
	mutex   sync.RWMutex
}

func (c *LockedMapStore[V]) Get(key int64) (V, bool) {

	c.mutex.Lock()
	value, ok := c.hashMap[key]
	c.mutex.Unlock()

	return value, ok
}

func (c *LockedMapStore[V]) Put(key int64, value V) {

	c.mutex.Lock()
	c.hashMap[key] = value
	c.mutex.Unlock()
}
func (c *LockedMapStore[V]) Delete(key int64) {

	c.mutex.Lock()
	delete(c.hashMap, key)
	c.mutex.Unlock()

}

func NewLockedMapStore[V any]() entities.Store[V] {
	store := &LockedMapStore[V]{
		hashMap: make(map[int64]V),
	}

	return store
}
