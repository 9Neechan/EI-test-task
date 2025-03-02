package db_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	sqlc "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
	util "github.com/9Neechan/EI-test-task/stats-service/internal/util"
)

// createRandomUser создает случайного пользователя для тестирования
func createRandomUser(t *testing.T) sqlc.User {
	name := util.RandomName()
	user, err := testQueries.CreateUser(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, name, user.Name)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

// TestCreateUser тестирует создание пользователя
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
