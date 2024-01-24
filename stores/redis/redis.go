package redis

import (
	"context"
	"runtime"
	"sync"

	"github.com/redis/go-redis/v9"
)

var once sync.Once

type Redis struct {
	Client *redis.Client
}

type Option func(*Redis)

func NewRedis(c Config, opts ...Option) (*Redis, error) {
	err := c.initConfig()
	if err != nil {
		return nil, err
	}

	if c.PoolSize == 0 {
		c.PoolSize = 10 * runtime.GOMAXPROCS(0)
	}

	r := &Redis{}

	for _, opt := range opts {
		opt(r)
	}

	r, err = newRedis(c, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func newRedis(c Config, r *Redis) (*Redis, error) {
	var err error
	once.Do(func() {
		r.Client = redis.NewClient(&redis.Options{
			Addr:     c.address(),
			Password: c.Password,
			DB:       c.DB,
			Protocol: c.Protocol,
			PoolSize: c.PoolSize,
		})

		_, err = r.Client.Ping(context.Background()).Result()
	})

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}