package grpc

import (
	"context"

	//"github.com/olezhek28/clean-architecture/internal/converter"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateService(ctx context.Context, req *desc.CreateServiceRequest) (*desc.CreateServiceResponse, error) {
	// Проверяем, что запрос не nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "❌ запрос не должен быть nil")
	}

	// Проверяем, что обязательные параметры заполнены
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "❌ имя сервиса не может быть пустым")
	}
	if req.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "❌ описание сервиса не может быть пустым")
	}
	if req.Price == 0 {
		return nil, status.Error(codes.InvalidArgument, "❌ цена сервиса не может быть пустым")
	}

	// Создаем сервис в БД
	arg := db.CreateServiceParams{
		Name:        req.Name,
		Description: req.Description,
		Price: req.Price,
	}

	service, err := i.db.CreateService(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "❌ ошибка создания сервиса: %v", err)
	}

	// Возвращаем успешный ответ
	return &desc.CreateServiceResponse{
		ServiceId:   service.ID,
		Name:        service.Name,
		Description: service.Description,
		Price: service.Price,
	}, nil
}
