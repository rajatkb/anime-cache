package services

import (
	"animecache/entities"

	"github.com/alphadose/haxmap"
)

type haxMapStore[V any] struct {
	store *haxmap.Map[int64, V]
}

func (c *haxMapStore[V]) Get(key int64) (V, bool) {
	value, ok := c.store.Get(key)
	return value, ok
}

func (c *haxMapStore[V]) Put(key int64, value V) {
	c.store.Set(key, value)
}
func (c *haxMapStore[V]) Delete(key int64) {
	c.store.Del(key)
}

func NewHaxMapStore[V any](size int) entities.Store[V] {
	store := &haxMapStore[V]{
		store: haxmap.New[int64, V](uintptr(size)),
	}

	return store
}
