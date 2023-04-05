package redisdb

import "github.com/redis/go-redis/v9"

type DB struct {
	*redis.Client
}

// New creates Redis client instance.
func New(uri string, password string) *DB {
	opts := &redis.Options{
		Addr: uri,
	}
	if password != "" {
		opts.Password = password
	}

	rdb := redis.NewClient(opts)

	return &DB{
		rdb,
	}
}
