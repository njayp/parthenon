package server

import (
	"context"
	"log"

	bff "github.com/njayp/parthenon/pkg/api"
)

type Server struct {
	bff.UnimplementedBFFServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) BoyfriendBot(context.Context, *bff.BoyfriendRequest) (*bff.BoyfriendResponse, error) {
	log.Print("Recieved Meow")
	return &bff.BoyfriendResponse{
		Emoji: "🐢",
	}, nil
}
