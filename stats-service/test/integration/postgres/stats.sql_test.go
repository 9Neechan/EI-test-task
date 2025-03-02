package db_test

import (
	"context"
	"testing"
	"time"

	sqlc "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
	"github.com/stretchr/testify/require"
)

// truncateTables очищает таблицу stats перед тестами
func truncateTables(t *testing.T) {
	_, err := testDB.Exec("TRUNCATE TABLE stats RESTART IDENTITY CASCADE")
	require.NoError(t, err)
}

// TestPostCall проверяет вставку данных в stats
func TestPostCall(t *testing.T) {
	//truncateTables(t) // Очистка перед тестом

	ctx := context.Background()
	arg := sqlc.PostCallParams{
		UserID:    1,
		ServiceID: 2,
	}

	stat, err := testQueries.PostCall(ctx, arg)
	require.NoError(t, err)
	require.Equal(t, arg.UserID, stat.UserID)
	require.Equal(t, arg.ServiceID, stat.ServiceID)
	//require.Equal(t, int64(1), stat.Count)
	require.WithinDuration(t, time.Now(), stat.CreatedAt, time.Second)
}

// TestGetCalls проверяет получение данных из stats
// ! TODO: дописать 
func TestGetCalls(t *testing.T) {
	//truncateTables(t)

	ctx := context.Background()
	arg := sqlc.PostCallParams{
		UserID:    1,
		ServiceID: 2,
	}

	// Вставляем тестовые данные
	_, err := testQueries.PostCall(ctx, arg)
	require.NoError(t, err)

	params := sqlc.GetStatsParams{
		UserID:    1,
		ServiceID: 0,
		Limit:     10,
		Offset:    0,
	}

	stats, err := testQueries.GetStats(ctx, params)
	require.NoError(t, err)
	require.Len(t, stats, 2)

	//require.Equal(t, int64(1), stats[0].UserID)
	//require.Equal(t, int64(2), stats[0].ServiceID)
	//require.Equal(t, int64(1), stats[0].Count)
}
