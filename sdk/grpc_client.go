package sdk

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCClient is a client for interacting with the service gRPC API
type GRPCClient struct {
	conn *grpc.ClientConn
	// Add your gRPC client stubs here
	// Example:
	// client pb.YourServiceClient
}

// GRPCClientOption is a functional option for configuring the gRPC client
type GRPCClientOption func(*grpcClientConfig)

type grpcClientConfig struct {
	dialOptions []grpc.DialOption
	timeout     time.Duration
}

// WithDialOptions allows setting custom gRPC dial options
func WithDialOptions(opts ...grpc.DialOption) GRPCClientOption {
	return func(c *grpcClientConfig) {
		c.dialOptions = append(c.dialOptions, opts...)
	}
}

// WithTimeout sets the default timeout for gRPC calls
func WithTimeout(timeout time.Duration) GRPCClientOption {
	return func(c *grpcClientConfig) {
		c.timeout = timeout
	}
}

// NewGRPCClient creates a new gRPC client
func NewGRPCClient(address string, opts ...GRPCClientOption) (*GRPCClient, error) {
	config := &grpcClientConfig{
		timeout:     30 * time.Second,
		dialOptions: []grpc.DialOption{},
	}

	for _, opt := range opts {
		opt(config)
	}

	// Add default dial options if none provided
	if len(config.dialOptions) == 0 {
		config.dialOptions = append(config.dialOptions,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
	}

	conn, err := grpc.NewClient(address, config.dialOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	return &GRPCClient{
		conn: conn,
		// Initialize your client stubs here
		// Example:
		// client: pb.NewYourServiceClient(conn),
	}, nil
}

// Close closes the gRPC connection
func (c *GRPCClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Add your gRPC client methods here
// Example:
// func (c *GRPCClient) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
//     return c.client.GetUser(ctx, req)
// }
