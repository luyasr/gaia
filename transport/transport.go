package transport

import (
	"context"
	"net"
	"strconv"
)

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}

func IsValidAddress(address string) bool {
	_, port, err := net.SplitHostPort(address)

	return err == nil && isValidPort(port)
}

func isValidPort(port string) bool {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return false
	}

	return portInt > 0 && portInt <= 65535
}
