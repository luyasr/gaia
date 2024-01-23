package redis

import (
	"context"
	"sync"

	"github.com/cenkalti/backoff/v4"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/reflection"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
	config Config
	once sync.Once
}

type Option func(*Redis)

func getClient(ctx context.Context, r *Redis) (*redis.Client, error) {
	var err error
	r.once.Do(func() {
		r.Client = redis.NewClient(&redis.Options{
			Addr:     r.config.Address(),
			Password: r.config.Password,
			DB:       r.config.DB,
			Protocol: r.config.Protocol,
		})

		_, err = r.Client.Ping(ctx).Result()
	})

	return r.Client, err
}

func (r *Redis) Reconnect(ctx context.Context) {
	go func() {
		for {
			if r.Client == nil {
				return
			}

			operation := func() error {
				_, err := r.Client.Ping(ctx).Result()
				return err
			}

			// 使用 backoff 库进行指数回避重连
			expBackOff := backoff.NewExponentialBackOff()
			backOffCtx := backoff.WithContext(expBackOff, ctx)
			if err := backoff.Retry(operation, backOffCtx); err != nil {
				log.Errorf("redis reconnect failed: %v", err)
			}
		}
	}()
}

func NewRedis(c Config, opts ...Option) (*Redis, error) {
	// 设置默认值
	reflection.SetUp(&c)
	// 校验 redis 配置信息
	if err := c.Validate(); err != nil {
		return nil, err
	}

	rds := &Redis{
		config: c,
	}

	for _, opt := range opts {
		opt(rds)
	}

	return newRedis(context.Background(), rds)
}

func newRedis(ctx context.Context, r *Redis) (*Redis, error) {
	var err error
	r.Client, err = getClient(ctx, r)
	if err != nil {
		log.Fatalf("redis connect failed: %v", err)
		return nil, err
	}

	r.Reconnect(ctx)

	return r, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}
