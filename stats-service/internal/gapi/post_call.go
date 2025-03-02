package grpc

import (
	"context"

	//"github.com/olezhek28/clean-architecture/internal/converter"
	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) PostCall(ctx context.Context, req *desc.PostCallRequest) (*desc.PostCallResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	// Проверяем, передал ли клиент обязательные параметры
	if req.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "UserId is required and must be greater than 0")
	}

	if req.ServiceId == 0 {
		return nil, status.Error(codes.InvalidArgument, "ServiceId is required and must be greater than 0")
	}

	arg := db.PostCallParams{
		UserID:    req.UserId,
		ServiceID: req.ServiceId,
	}

	_, err := i.db.PostCall(ctx, arg)
	if err != nil {
		return &desc.PostCallResponse{
			Success: false,
		}, status.Errorf(codes.Internal, "failed to post call: %v", err)
	}

	return &desc.PostCallResponse{
		Success: true,
	}, nil
}
