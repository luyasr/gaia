package transport

import "context"

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}
