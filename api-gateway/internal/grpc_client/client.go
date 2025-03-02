package grpcclient

import (
	desc "github.com/9Neechan/EI-test-task/api/pb"
	"google.golang.org/grpc"
)

// GRPCClient - a wrapper for the gRPC client
type GRPCClient struct {
	conn            *grpc.ClientConn
	statsClient     desc.StatsServiceClient
}

// NewGRPCClient creates a new gRPC client
func NewGRPCClient(conn *grpc.ClientConn) *GRPCClient {
	return &GRPCClient{
		conn:        conn,
		statsClient: desc.NewStatsServiceClient(conn),
	}
}

