package redis

import (
	"context"
	"runtime"
	"sync"

	"github.com/luyasr/gaia/reflection"
	"github.com/redis/go-redis/v9"
)

var once sync.Once

type Redis struct {
	Client *redis.Client
	config Config
}

type Option func(*Redis)

func NewRedis(c Config, opts ...Option) (*Redis, error) {
	reflection.SetUp(&c)

	if c.PoolSize == 0 {
		c.PoolSize = 10 * runtime.GOMAXPROCS(0)
	}

	rds, err := newRedis(c, opts...)
	if err != nil {
		return nil, err
	}

	return rds, nil
}

func newRedis(c Config, opts ...Option) (*Redis, error) {
	var err error

	r := &Redis{
		config: c,
	}

	for _, opt := range opts {
		opt(r)
	}

	once.Do(func() {
		r.Client = redis.NewClient(&redis.Options{
			Addr:     r.config.Address(),
			Password: r.config.Password,
			DB:       r.config.DB,
			Protocol: r.config.Protocol,
			PoolSize: r.config.PoolSize,
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
