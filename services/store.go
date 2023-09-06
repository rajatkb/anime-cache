package services

import (
	"animecache/entities"
)

var GET_OP int8 = 4
var PUT_OP int8 = 2
var DEL_OP int8 = 3

type Ops[V any] struct {
	Op   int8
	Key  int64
	Data V
	Resp chan V
}

type Store[V any] struct {
	hashMap          map[int64]V
	operationChannel chan Ops[V]
}

func (c *Store[V]) Get(key int64) (V, bool) {

	resp := make(chan V)
	c.operationChannel <- Ops[V]{
		Op:   GET_OP,
		Key:  key,
		Resp: resp,
	}

	value, ok := <-resp

	return value, ok
}

func (c *Store[V]) Put(key int64, value V) {
	c.operationChannel <- Ops[V]{
		Op:   PUT_OP,
		Key:  key,
		Data: value,
	}
}
func (c *Store[V]) Delete(key int64) {
	c.operationChannel <- Ops[V]{
		Op:  PUT_OP,
		Key: key,
	}
}

func NewStore[V any](operationBufferSize int) entities.Store[V] {
	store := &Store[V]{
		hashMap:          make(map[int64]V),
		operationChannel: make(chan Ops[V], operationBufferSize),
	}
	go func() {

		// runtime.LockOSThread()

		for ops := range store.operationChannel {
			switch ops.Op {
			case PUT_OP:
				store.hashMap[ops.Key] = ops.Data
			case DEL_OP:
				delete(store.hashMap, ops.Key)
			case GET_OP:
				if value, ok := store.hashMap[ops.Key]; ok {
					ops.Resp <- value
				} else {
					close(ops.Resp)
				}
			}
		}
	}()

	return store
}
