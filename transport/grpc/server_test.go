package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServerRun(t *testing.T) {
	s := NewServer(Address("invalid"))

	err := s.Run()
	assert.Error(t, err)
	assert.Equal(t, defaultAddress, s.address)
}

func TestServerShutdown(t *testing.T) {
	s := NewServer()

	go func() {
		_ = s.Run()
	}()

	time.Sleep(3 * time.Second)

	err := s.Shutdown(context.Background())
	assert.NoError(t, err)
}
