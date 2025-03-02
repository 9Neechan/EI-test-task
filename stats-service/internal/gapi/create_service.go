package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

// CreateService creates a new service based on the request parameters.
func (i *Implementation) CreateService(ctx context.Context, req *desc.CreateServiceRequest) (*desc.CreateServiceResponse, error) {
	// Validate the request to ensure it is not nil.
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request must not be nil")
	}

	// Validate the required parameters to ensure they are not empty.
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "service name cannot be empty")
	}
	if req.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "service description cannot be empty")
	}
	if req.Price == 0.0 {
		return nil, status.Error(codes.InvalidArgument, "service price cannot be empty")
	}

	arg := db.CreateServiceParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	service, err := i.db.CreateService(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating service: %v", err)
	}

	return &desc.CreateServiceResponse{
		ServiceId:   service.ID,
		Name:        service.Name,
		Description: service.Description,
		Price:       service.Price,
	}, nil
}
