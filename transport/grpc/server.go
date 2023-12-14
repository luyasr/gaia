package grpc

import (
	"context"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	server  *grpc.Server
	log     *log.Helper
	address string
}

type ServerOption func(*Server)

func Address(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}

func NewServer(opt ...ServerOption) *Server {
	svc := &Server{
		server: &grpc.Server{},
		log:    log.NewHelper(zerolog.New(zerolog.NewConsoleLogger())),
	}

	for _, o := range opt {
		o(svc)
	}

	return svc
}

func (s *Server) Run() error {
	s.log.Infof("grpc server listen on %s", s.address)
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		s.log.Fatalf("grpc server run fatal, error: %s", err)
	}

	return s.server.Serve(listen)
}

func (s *Server) Shutdown(context.Context) error {
	s.log.Info("grpc server shutdown...")
	s.server.GracefulStop()

	return nil
}
