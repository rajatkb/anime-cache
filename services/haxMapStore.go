package services

// type haxMapStore[V any] struct {
// 	store
// }

// func (c *haxMapStore[V]) Get(key int64) (V, bool) {

// 	c.mutex.Lock()
// 	value, ok := c.hashMap[key]
// 	c.mutex.Unlock()

// 	return value, ok
// }

// func (c *haxMapStore[V]) Put(key int64, value V) {

// 	c.mutex.Lock()
// 	c.hashMap[key] = value
// 	c.mutex.Unlock()
// }
// func (c *haxMapStore[V]) Delete(key int64) {

// 	c.mutex.Lock()
// 	delete(c.hashMap, key)
// 	c.mutex.Unlock()

// }

// func NewHaxMapStore[V any](operationBufferSize int, lockOsThreads bool) entities.Store[V] {
// 	store := &LockedMapStore[V]{
// 		hashMap: make(map[int64]V),
// 	}

// 	return store
// }
