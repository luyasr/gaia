package mongo

import (
	"context"
	"time"

	"github.com/luyasr/gaia/errors"
	"github.com/luyasr/gaia/ioc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const Name = "mongo"

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
		return errors.Internal("mongo type assertion failed", "expected *Config, got %T", cfg)
	}

	rdb, err := New(mongoCfg)
	if err != nil {
		return err
	}
	m.Client = rdb.Client

	return nil
}

func (m *Mongo) Name() string {
	return Name
}

func New(c *Config, opts ...Option) (*Mongo, error) {
	cfg, err := c.initConfig()
	if err != nil {
		return nil, err
	}

	m := &Mongo{}

	for _, opt := range opts {
		opt(m)
	}

	return new(cfg, m)
}

func new(c *Config, m *Mongo) (*Mongo, error) {
	var err error
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(c.uri()).SetServerAPIOptions(serverAPI)

	m.Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	err = m.ping()
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
