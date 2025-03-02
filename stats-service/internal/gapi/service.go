package grpc

import (
	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

// Implementation represents the service implementation
type Implementation struct {
	db db.Querier
	desc.UnimplementedStatsServiceServer
}

// NewImplementation creates a new service implementation
func NewImplementation(repo db.Querier) *Implementation {
	return &Implementation{db: repo}
}
