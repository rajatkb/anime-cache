package animecache

import (
	"animecache/entities"
	"animecache/services"
)

type Cache[V any] struct {
	config CacheConfig
	stores []services.LockedMapStore[V]
}

func (c *Cache[V]) Get(key []byte) (V, bool) {
	var empty V
	return empty, false
}

func (c *Cache[V]) Put(key []byte, value V) {

}
func (c *Cache[V]) Delete(key []byte) {

}

func NewCache[V any](config CacheConfig) entities.AnimeCache[V] {
	cache := &Cache[V]{
		config: config,
		stores: make([]services.LockedMapStore[V], config.MaxShards),
	}

	return cache
}
