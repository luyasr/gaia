package http

import (
	"context"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
	"github.com/luyasr/gaia/transport"
	"net/http"
)

var _ transport.Server = (*Server)(nil)

type Server struct {
	server *http.Server
	log    *log.Helper
}

type ServerOption func(*Server)

func WithAddr(addr string) ServerOption {
	return func(s *Server) {
		s.server.Addr = addr
	}
}

func WithHandler(handler http.Handler) ServerOption {
	return func(s *Server) {
		s.server.Handler = handler
	}
}

func NewServer(opts ...ServerOption) *Server {
	logger := zerolog.New(zerolog.NewConsoleLogger())

	svc := &Server{
		server: &http.Server{},
		log:    log.NewHelper(logger),
	}

	for _, opt := range opts {
		opt(svc)
	}

	return svc
}

func (s *Server) Run() error {
	s.log.Infof("http server listen on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Infof("http server shutdown...")
	return s.Shutdown(ctx)
}
