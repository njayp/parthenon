package main

import (
	"flag"

	"github.com/njayp/parthenon/pkg/daemon"
	"github.com/njayp/parthenon/pkg/server/grpcServer"
	"github.com/njayp/parthenon/pkg/server/httpServer"

	"k8s.io/klog/v2"

	"google.golang.org/grpc"
)

var (
	/*
		tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
		certFile   = flag.String("cert_file", "", "The TLS cert file")
		keyFile    = flag.String("key_file", "", "The TLS key file")
		jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	*/
	grpcPort = flag.Int("gport", 9090, "The grpc port")
	httpPort = flag.Int("hport", 8080, "The http port")
)

func main() {
	err := daemonMain()
	if err != nil {
		klog.Fatal(err)
	}
}

func daemonMain() error {
	d := daemon.NewDaemon("parth")
	return d.Stop()
}

func serverMain() error {
	flag.Parse()
	ch := make(chan error)
	go func() {
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
					klog.Fatalf("Failed to generate credentials %v", err)
				}
				opts = []grpc.ServerOption{grpc.Creds(creds)}
			}
		*/
		ch <- grpcServer.Start(*grpcPort, opts)
	}()

	go func() {
		ch <- httpServer.Start(*httpPort)
	}()

	return <-ch
}
