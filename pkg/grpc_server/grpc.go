package grpc_server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	server *grpc.Server
	port   int
	host   string
}

func New(port int, host string) *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
		port:   port,
		host:   host,
	}
}

func (g *GRPCServer) Run() {

	l, err := net.Listen("tcp", fmt.Sprintf("%v:%v", g.host, g.port))
	if err != nil {
		panic(err)
	}

	if err = g.server.Serve(l); err != nil {
		panic(err)
	}
}

func (g *GRPCServer) GracefulStop() {
	g.server.GracefulStop()
}

func (g *GRPCServer) RegisterService(desc *grpc.ServiceDesc, impl any) {
	g.server.RegisterService(desc, impl)
}
