package grpcclient

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/9Neechan/EI-test-task/api/pb"
)

// GetStats calls the GetStats gRPC method on the server
func (c *GRPCClient) GetStats(ctx context.Context, in *desc.GetStatsRequest) (*desc.GetStatsResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := c.statsClient.GetStats(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("Error calling GetStats: %w", err)
	}

	return res, nil
}
