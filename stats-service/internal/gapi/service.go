package grpc

import (
	//"github.com/9Neechan/EI-test-task/stats-service/internal/service"

	"fmt"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	db "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
)

type Implementation struct {
	db db.Querier
	desc.UnimplementedStatsServiceServer
	//statsService service.StatsService
	
}

func NewImplementation(repo db.Querier) *Implementation {
impl := &Implementation{db: repo}

    fmt.Printf("✅ Created Implementation: %+v\n", impl) // Вывод всей структуры
    return impl
}

//cd stats-service/cmd/grpc_server/

