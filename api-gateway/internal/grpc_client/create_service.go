package grpcclient

import (
	"context"
	"time"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *GRPCClient) CreateService(ctx context.Context, in *desc.CreateServiceRequest) (*desc.CreateServiceResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := c.statsClient.CreateService(ctx, in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create service: %v", err)
	}

	return resp, nil
}
