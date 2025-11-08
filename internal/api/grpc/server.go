package grpc

import (
	"fmt"
	"net"

	"github.com/travisbale/go-template/internal/db/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Config struct {
	Address string
	DB      *postgres.DB
}

// Server implements the gRPC service
type Server struct {
	Addr string
	*grpc.Server
}

// NewServer creates a new gRPC server
func NewServer(config *Config) *Server {
	grpcServer := grpc.NewServer()

	// Enable gRPC reflection for development/debugging with grpcurl
	reflection.Register(grpcServer)

	// Register your gRPC services here
	// Example:
	// pb.RegisterYourServiceServer(grpcServer, yourServiceHandler)

	return &Server{
		Addr:   config.Address,
		Server: grpcServer,
	}
}

func (s *Server) ListenAndServe() error {
	// Create gRPC listener
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("failed to create gRPC listener: %w", err)
	}

	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("gRPC server error: %w", err)
	}

	return nil
}
