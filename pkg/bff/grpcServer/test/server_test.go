package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/njayp/parthenon/pkg/api"
	"github.com/njayp/parthenon/pkg/bff/grpcServer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPC(t *testing.T) {
	port := 9090
	go grpcServer.Start(port, nil)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf(":%v", port), opts...)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer conn.Close()
	client := api.NewBFFClient(conn)

	t.Run("meow", func(t *testing.T) {
		response, err := client.BoyfriendBot(context.TODO(), &api.BoyfriendRequest{})
		if err != nil {
			t.Fatal(err.Error())
		}
		text := response.GetEmoji()
		expected := "🐢"
		if text != expected {
			t.Errorf("response text '%s' does not contian '%s'", text, expected)
		}
	})
}
