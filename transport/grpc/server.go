package grpc

import (
	"context"
	"net"

	"github.com/luyasr/gaia/log"
	"github.com/luyasr/gaia/log/zerolog"
	"github.com/luyasr/gaia/transport"
	"google.golang.org/grpc"
)

const (
	defaultAddress = ":50051"
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
		server: grpc.NewServer(),
		log:    log.NewHelper(zerolog.New(zerolog.DefaultLogger)),
	}

	for _, o := range opt {
		o(svc)
	}

	return svc
}

func (s *Server) Run() error {
	if !transport.IsValidAddress(s.address) {
		s.log.Warnf("grpc server address %s is invalid, use default address %s", s.address, defaultAddress)
		s.address = defaultAddress
	}

	s.log.Infof("grpc server listen on %s", s.address)
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		s.log.Fatalf("grpc server run fatal, error: %s", err)
	}

	return s.server.Serve(listen)
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("grpc server shutdown...")
	s.server.GracefulStop()

	return nil
}
