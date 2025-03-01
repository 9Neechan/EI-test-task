package service

import (
	"context"

	"github.com/9Neechan/EI-test-task/stats-service/internal/model"
)

type StatsService interface {
	CreateService(ctx context.Context, in *model.CreateServiceRequest) (*model.CreateServiceResponse, error)
	PostCall(ctx context.Context, in *model.PostCallRequest) (bool, error)
	GetStats(ctx context.Context, in *model.GetStatsRequest) (*model.GetStatsResponse, error)
}
