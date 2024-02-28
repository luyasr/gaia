package mongo

import (
	"github.com/luyasr/gaia/ioc"
	"go.mongodb.org/mongo-driver/mongo"
)

func Client() *mongo.Client {
	return ioc.Container.Get(ioc.DbNamespace, Name).(*Mongo).Client
}