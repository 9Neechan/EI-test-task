package grpcclient

import (
	"context"
	desc "github.com/9Neechan/EI-test-task/api/pb"
)

func (c *GRPCClient) PostCall(ctx context.Context, in *desc.PostCallRequest) (*desc.PostCallResponse, error) {
	return nil, nil
}