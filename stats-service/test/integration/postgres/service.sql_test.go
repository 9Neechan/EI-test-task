package db_test

import (
	"context"
	"testing"

	sqlc "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
	util "github.com/9Neechan/EI-test-task/stats-service/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomService(t *testing.T) sqlc.Service {
	arg := sqlc.CreateServiceParams{
		Name:        util.RandomName(),
		Description: util.RandomDescription(),
	}

	service, err := testQueries.CreateService(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, service)

	require.Equal(t, arg.Name, service.Name)
	require.Equal(t, arg.Description, service.Description)

	require.NotZero(t, service.ID)
	require.NotZero(t, service.CreatedAt)

	return service
}

func TestCreateService(t *testing.T) {
	createRandomService(t)
}