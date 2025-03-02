package grpc

import (
	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

type Implementation struct {
	db db.Querier
	desc.UnimplementedStatsServiceServer
}

func NewImplementation(repo db.Querier) *Implementation {
	return &Implementation{db: repo}
}
