package grpcclient

import (
	"context"
	"fmt"
	"time"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetStats вызывает gRPC-метод GetStats на сервере
func (c *GRPCClient) GetStats(ctx context.Context, in *desc.GetStatsRequest) (*desc.GetStatsResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := c.statsClient.GetStats(ctx, in)
	if err != nil {
		return nil, fmt.Errorf("❌ Ошибка вызова GetStats: %w", err)
	}

	return res, nil
}
