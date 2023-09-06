package animecache

type CacheConfig struct {
	MaxShards   int // each shard represents a single cache store. more the shards lesser the cache contention
	MaxShardLen int // maximum items per shard
}
