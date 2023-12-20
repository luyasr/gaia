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
		log:    log.NewHelper(zerolog.New(zerolog.DefaultLogger)),
	}

	for _, o := range opt {
		o(svc)
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
