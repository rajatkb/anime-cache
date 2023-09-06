package entities

type AnimeCache[V any] interface {
	/*
		Get value
	*/
	Get(key []byte) (V, bool)
	/*
		Put value with lifetime till evicted
	*/
	Put(key []byte, value V)
	/*
		Delete value
	*/
	Delete(key []byte)

	// PutTTL(key []byte, value V, sec time.Duration)
}

type Store[V any] interface {
	/*
		Get value
	*/
	Get(key int64) (V, bool)
	/*
		Put value with lifetime till evicted
	*/
	Put(key int64, value V)
	/*
		Delete value
	*/
	Delete(key int64)

	// PutTTL(key []byte, value V, sec time.Duration)
}
