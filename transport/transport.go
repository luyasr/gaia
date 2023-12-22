package transport

import (
	"context"
	"net"
)

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}

func IsValidAddress(address string) bool {
	_, port, err := net.SplitHostPort(address)

	return err == nil && port != "0"
}
