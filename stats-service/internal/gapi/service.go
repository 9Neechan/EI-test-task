package grpc

import (
	//"github.com/9Neechan/EI-test-task/stats-service/internal/service"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

type Implementation struct {
	db db.Querier
	desc.UnimplementedStatsServiceServer
	//statsService service.StatsService
}

func NewImplementation(repo db.Querier) *Implementation {
	return &Implementation{db: repo}
}

//cd stats-service/cmd/grpc_server/
