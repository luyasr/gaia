package app

import (
	"github.com/luyasr/gaia/transport/grpc"
	"github.com/luyasr/gaia/transport/http"
	"testing"
)

func TestNew(t *testing.T) {
	httpServer := http.NewServer(http.Address(""))
	grpcServer := grpc.NewServer(grpc.Address(""))
	app := New(Server(httpServer, grpcServer))
	err := app.Run()
	t.Fatal(err)
}
