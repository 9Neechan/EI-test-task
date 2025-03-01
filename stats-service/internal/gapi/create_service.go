package grpc

import (
	"context"

	//"github.com/olezhek28/clean-architecture/internal/converter"
	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

func (i *Implementation) CreateService(ctx context.Context, req *desc.CreateServiceRequest) (*desc.CreateServiceResponse, error) {
	//uuid, err := i.statsService.CreateService(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	arg := db.CreateServiceParams{
		Name: req.Name,
		Description: req.Description,
	}

	service, err := i.db.CreateService(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &desc.CreateServiceResponse{
		ServiceId:   service.ID,
		Name:        service.Name,
		Description: service.Description,
	}, nil
}
