package redis

import (
	"github.com/luyasr/gaia/ioc"
	"github.com/redis/go-redis/v9"
)

func Client() *redis.Client {
	return ioc.Container.Get(ioc.DbNamespace, Name).(*Redis).Client
}