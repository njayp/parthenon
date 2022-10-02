package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/njayp/parthenon/pkg/api"
	"google.golang.org/grpc"
)

type Server struct {
	BFFServer
}

func Start(port int, opts []grpc.ServerOption) error {
	grpcServer := grpc.NewServer(opts...)
	server := &Server{}
	api.RegisterBFFServer(grpcServer, server)
	return grpcServer.Serve(makeListener(port))
}

func makeListener(port int) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("grpc listening on port %v", port)
	return lis
}
