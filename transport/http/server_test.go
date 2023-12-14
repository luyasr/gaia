package http

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestServerRun(t *testing.T) {
	s := NewServer(Address(":8080"), Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))

	go func() {
		err := s.Run()
		assert.NoError(t, err)
	}()

	time.Sleep(time.Second) // give server time to start

	resp, err := http.Get("http://localhost:8080")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestServerShutdown(t *testing.T) {
	s := NewServer(Address(":8080"), Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))

	go func() {
		err := s.Run()
		assert.NoError(t, err)
	}()

	time.Sleep(time.Second) // give server time to start

	err := s.Shutdown(context.Background())
	assert.NoError(t, err)

	resp, err := http.Get("http://localhost:8080")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestServerWithInvalidAddr(t *testing.T) {
	s := NewServer(Address("invalid"), Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})))

	err := s.Run()
	assert.Error(t, err)
}

func TestServerWithNilHandler(t *testing.T) {
	s := NewServer(Address(":8080"), Handler(nil))

	err := s.Run()
	assert.Error(t, err)
}
