package redis

import (
	"context"
	"testing"
)

func TestNewRedis(t *testing.T) {
	config := Config{
		Password: "12345678",
	}

	r, err := NewRedis(config)
	if err != nil {
		t.Fatal(err)
	}

	result, err := r.Client.Ping(context.Background()).Result()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
