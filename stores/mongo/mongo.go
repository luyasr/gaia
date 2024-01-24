package mongo

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var once sync.Once

type Mongo struct {
	Client *mongo.Client
}

type Option func(*Mongo)

func NewMongo(c Config, opts ...Option) (*Mongo, error) {
	err := c.initConfig()
	if err != nil {
		return nil, err
	}

	m := &Mongo{}

	for _, opt := range opts {
		opt(m)
	}

	m, err = newMongo(c, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func newMongo(c Config, m *Mongo) (*Mongo, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(c.uri()).SetServerAPIOptions(serverAPI)

	var err error
	once.Do(func() {
		m.Client, err = mongo.Connect(context.TODO(), opts)

		err = m.ping()
	})
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Mongo) Close() error {
	if m.Client != nil {
		return m.Client.Disconnect(context.TODO())
	}

	return nil
}

func (m *Mongo) ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := m.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}
