package grpc

import (
	"context"

	//"github.com/olezhek28/clean-architecture/internal/converter"
	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetStats(ctx context.Context, req *desc.GetStatsRequest) (*desc.GetStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	arg := db.GetStatsParams{}

	// Проверяем указатели, чтобы избежать разыменования nil
	if req.UserId != nil {
		arg.UserID = *req.UserId
	}

	if req.ServiceId != nil {
		arg.ServiceID = *req.ServiceId
	}

	arg.Limit = req.Limit
	arg.Offset = req.Page

	stats_db, err := i.db.GetStats(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get stats: %v", err)
	}

	// Предварительное выделение памяти для слайса
	stats := make([]*desc.StatRecord, 0, len(stats_db))
	for _, val := range stats_db {
		stats = append(stats, &desc.StatRecord{
			UserId:    val.UserID,
			ServiceId: val.ServiceID,
			Count:     val.Count,
		})
	}

	return &desc.GetStatsResponse{Stats: stats}, nil
}
