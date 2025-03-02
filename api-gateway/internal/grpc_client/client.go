package grpcclient

import (
	desc "github.com/9Neechan/EI-test-task/api/pb"
	"google.golang.org/grpc"
)

// GRPCClient - обёртка для gRPC-клиента
type GRPCClient struct {
	conn            *grpc.ClientConn
	statsClient     desc.StatsServiceClient
}

// NewGRPCClient создаёт новый gRPC-клиент
func NewGRPCClient(conn *grpc.ClientConn) *GRPCClient {
	return &GRPCClient{
		conn:        conn,
		statsClient: desc.NewStatsServiceClient(conn),
	}
}


