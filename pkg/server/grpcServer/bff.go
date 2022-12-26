package grpcServer

import (
	"context"

	"github.com/njayp/parthenon/pkg/api"
	"k8s.io/klog/v2"
)

type BFFServer struct {
	api.UnimplementedBFFServer
}

func (s *BFFServer) BoyfriendBot(context.Context, *api.BoyfriendRequest) (*api.BoyfriendResponse, error) {
	klog.Infof("Recieved Meow")
	return &api.BoyfriendResponse{
		Emoji: "🐢",
	}, nil
}
