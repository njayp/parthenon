package grpcServer

import (
	"fmt"
	"net"

	"github.com/njayp/parthenon/pkg/api"
	"google.golang.org/grpc"
	"k8s.io/klog/v2"
)

type Server struct {
	BFFServer
}

func Start(port int, opts []grpc.ServerOption) error {
	klog.Infof("grpc listening on port %v", port)
	address := fmt.Sprintf(":%v", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(opts...)
	server := &Server{}
	api.RegisterBFFServer(grpcServer, server)
	return grpcServer.Serve(lis)
}
