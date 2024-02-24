package http

import (
	"context"
	"net/http"

	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
	"github.com/luyasr/gaia/transport"
)

var _ transport.Server = (*Server)(nil)

const (
	// defaultAddress is the default address of http server
	defaultAddress = ":8080"
)

type Server struct {
	server *http.Server
	log    *log.Helper
}

type ServerOption func(*Server)

func Address(address string) ServerOption {
	return func(s *Server) {
		s.server.Addr = address
	}
}

func Handler(handler http.Handler) ServerOption {
	return func(s *Server) {
		s.server.Handler = handler
	}
}

func Logger(logger *log.Helper) ServerOption {
	return func(s *Server) {
		s.log = logger
	}
}

func NewServer(opt ...ServerOption) *Server {
	svc := &Server{
		server: &http.Server{},
		log:    log.NewHelper(zerolog.DefaultLogger),
	}

	for _, o := range opt {
		o(svc)
	}

	return svc
}

func (s *Server) Run() error {
	if !transport.IsValidAddress(s.server.Addr) {
		s.log.Warnf("http server address %s is invalid, use default address %s", s.server.Addr, defaultAddress)
		s.server.Addr = defaultAddress
	}

	s.log.Infof("http server listen on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("http server shutdown...")
	return s.server.Shutdown(ctx)
}
