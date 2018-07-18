package server

import (
	"net"

	"github.com/tnngo/nog"
	"github.com/tnngo/ring/server/proto/provider"
	"google.golang.org/grpc"
)

type RingServer struct {
	IP   string
	Port int32
}

func New() *RingServer {
	return &RingServer{}
}

func ListenTCP() {
	tcpaddr, err := net.ResolveTCPAddr("tcp", ":8718")
	if err != nil {
		nog.SL.Error(err)
		return
	}

	l, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		nog.SL.Error(err)
		return
	}

	s := grpc.NewServer()
	provider.RegisterProviderServer(s, &Handle{})
	s.Serve(l)
}
