package mongo

import (
	"context"
	"sync"
	"time"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const name = "mongo"

var once sync.Once

type Mongo struct {
	Client *mongo.Client
}

type Option func(*Mongo)

func init() {
	ioc.Container.Registry(ioc.DbNamespace, &Mongo{})
}

func (m *Mongo) Init() error {
	cfg, ok := ioc.Container.GetFieldValueByConfig("Mongo")
	if !ok {
		return nil
	}

	mongoCfg, ok := cfg.(*Config)
	if !ok {
		return errors.Internal("mongo", "Mongo type assertion failed, expected *Config, got %T", cfg)
	}

	rdb, err := New(mongoCfg)
	if err != nil {
		return err
	}
	m.Client = rdb.Client

	return nil
}

func (m *Mongo) Name() string {
	return name
}

func New(c *Config, opts ...Option) (*Mongo, error) {
	if err := c.initConfig(); err != nil {
		return nil, err
	}

	m := &Mongo{}

	for _, opt := range opts {
		opt(m)
	}

	m, err := new(c, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func new(c *Config, m *Mongo) (*Mongo, error) {
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

	return m.Client.Ping(ctx, readpref.Primary())
}
