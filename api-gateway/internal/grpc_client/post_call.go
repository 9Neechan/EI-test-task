package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/9Neechan/EI-test-task/api/pb"
)

// PostCall calls the PostCall method on the server
func (c *GRPCClient) PostCall(ctx context.Context, in *desc.PostCallRequest) (*desc.PostCallResponse, error) {
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := c.statsClient.PostCall(ctx, in)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to post call: %v", err)
	}

	return resp, nil
}
