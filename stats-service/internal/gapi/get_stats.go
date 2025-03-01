package grpc

import (
	"context"

	//"github.com/olezhek28/clean-architecture/internal/converter"
	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

func (i *Implementation) GetStats(ctx context.Context, req *desc.GetStatsRequest) (*desc.GetStatsResponse, error) {
	arg := db.GetStatsParams{
		UserID:    *req.UserId,
		ServiceID: *req.ServiceId,
		Limit:     req.Limit,
		Offset:    req.Limit,
	}

	stats_db, err := i.db.GetStats(ctx, arg)
	if err != nil {
		return nil, err
	}

	// ! gorutines
	stats := make([]*desc.StatRecord, 0, len(stats_db))
	for _, val := range stats_db {
		stat_grpc := &desc.StatRecord{
			UserId:    val.UserID,
			ServiceId: val.ServiceID,
			Count:     val.Count,
		}
		stats = append(stats, stat_grpc)
	}

	return &desc.GetStatsResponse{
		Stats: stats,
	}, nil
}
