package grpc

import (
	"context"

	//"github.com/olezhek28/clean-architecture/internal/converter"
	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

func (i *Implementation)PostCall(ctx context.Context, req *desc.PostCallRequest) (*desc.PostCallResponse, error) {
	arg := db.PostCallParams{
		UserID: req.UserId,
		ServiceID: req.ServiceId,
	}

	_, err := i.db.PostCall(ctx, arg)
	if err != nil {
		return &desc.PostCallResponse{
			Success: false,
		}, err
	}

	return &desc.PostCallResponse{
		Success: true,
	}, nil
}
