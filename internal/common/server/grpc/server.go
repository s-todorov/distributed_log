package grpc

import (
	v4 "distributed_log/internal/common/api/protobuf/v4"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

// Server defines the available operations for gRPC server.
type Server interface {
	// Serve is called for serving requests.
	Serve() error
	// GracefulStop is called for stopping the server.
	GracefulStop()
}

// Timeout is the number of seconds to attempt a graceful shutdown, or
// for timing out read/write operations
const Timeout = 30 * time.Second

// ServicesConfiguration is an alias for a function that will take in a pointer to an Services and modify it
type ServicesConfiguration func(os *GrpcServer) error

// GrpcServer implements Server.
type GrpcServer struct {
	listener net.Listener
	Rpc      *grpc.Server
	opts     []grpc.ServerOption
	certFile string
	keyFile  string
}

func NewServer(port int, cfgs ...ServicesConfiguration) (*GrpcServer, error) {
	var opts []grpc.ServerOption
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return nil, err
	}
	os := &GrpcServer{
		Rpc:      grpc.NewServer(opts...),
		listener: lis,
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

func Register(srv v4.LogServer) ServicesConfiguration {
	return func(os *GrpcServer) error {
		v4.RegisterLogServer(os.Rpc, srv)
		return nil
	}
}

func (s *GrpcServer) Serve() error {
	log.Printf("server listening at %v", s.listener.Addr())
	return s.Rpc.Serve(s.listener)
}

func (s *GrpcServer) GracefulStop() {
	s.Rpc.GracefulStop()
}
