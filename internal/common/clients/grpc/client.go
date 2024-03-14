package grpc

import (
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// Timeout is the number of seconds to attempt a graceful shutdown, or
// for timing out read/write operations
const Timeout = 30 * time.Second

var serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")

// ServicesConfiguration is an alias for a function that will take in a pointer to an Services and modify it
type ServicesConfiguration func(os *GrpcClient) error

// GrpcServer implements Server.
type GrpcClient struct {
	clientConn    *grpc.ClientConn
	opts          []grpc.DialOption
	Client        v4.LogClient
	certFile      string
	keyFile       string
	serverAddress string
}

func NewClient(serverAddress string, cfgs ...ServicesConfiguration) (*GrpcClient, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	os := &GrpcClient{
		Client:        v4.NewLogClient(conn),
		clientConn:    conn,
		serverAddress: serverAddress,
	}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func (s *GrpcClient) Options() []grpc.DialOption {
	return s.opts
}

func (s *GrpcClient) GracefulStop() {
	defer s.clientConn.Close()
}
