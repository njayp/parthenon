package client

import (
	"context"
	"log"

	"github.com/njayp/parthenon/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Meow() error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("bff-svc.default:90", opts...)
	if err != nil {
		return err
	}
	defer conn.Close()
	client := api.NewBFFClient(conn)
	response, err := client.BoyfriendBot(context.TODO(), &api.BoyfriendRequest{})
	if err != nil {
		return err
	}
	log.Print(response.GetEmoji())
	return nil
}
