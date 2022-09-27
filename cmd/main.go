package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	bff "github.com/njayp/parthenon/pkg/api"
	"github.com/njayp/parthenon/pkg/bff/server"
	"google.golang.org/grpc"
)

var (
	/*
		tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
		certFile   = flag.String("cert_file", "", "The TLS cert file")
		keyFile    = flag.String("key_file", "", "The TLS key file")
		jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	*/
	port = flag.Int("port", 9090, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on port %v", *port)
	var opts []grpc.ServerOption
	/*
		if *tls {
			if *certFile == "" {
				*certFile = data.Path("x509/server_cert.pem")
			}
			if *keyFile == "" {
				*keyFile = data.Path("x509/server_key.pem")
			}
			creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
			if err != nil {
				log.Fatalf("Failed to generate credentials %v", err)
			}
			opts = []grpc.ServerOption{grpc.Creds(creds)}
		}
	*/

	grpcServer := grpc.NewServer(opts...)
	bff.RegisterBFFServer(grpcServer, server.NewServer())
	grpcServer.Serve(lis)
}
