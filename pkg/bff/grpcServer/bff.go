package grpcServer

import (
	"context"
	"log"

	"github.com/njayp/parthenon/pkg/api"
)

type BFFServer struct {
	api.UnimplementedBFFServer
}

func (s *BFFServer) BoyfriendBot(context.Context, *api.BoyfriendRequest) (*api.BoyfriendResponse, error) {
	log.Print("Recieved Meow")
	return &api.BoyfriendResponse{
		Emoji: "🐢",
	}, nil
}
