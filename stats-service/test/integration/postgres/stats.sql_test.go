package db_test

import (
	"context"
	"testing"

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

	service1 := createRandomService(t)
	serviceID1 := service1.ID

	user1 := createRandomUser(t)
	userID1 := user1.ID

	ctx := context.Background()
	arg := sqlc.PostCallParams{
		UserID:    userID1,
		ServiceID: serviceID1,
	}

	for i := 0; i < 3; i++ {
		stat, err := testQueries.PostCall(ctx, arg)
		require.NoError(t, err)
		require.Equal(t, arg.UserID, stat.UserID)
		require.Equal(t, arg.ServiceID, stat.ServiceID)
		require.Equal(t, int64(i+1), stat.Count)
	}
}

func TestPostCallNotExistingIDs(t *testing.T) {
	//truncateTables(t) // Очистка перед тестом

	ctx := context.Background()
	arg := sqlc.PostCallParams{
		UserID:    200,
		ServiceID: 1,
	}

	_, err := testQueries.PostCall(ctx, arg)
	require.Error(t, err)

	arg2 := sqlc.PostCallParams{
		UserID:    1,
		ServiceID: 200,
	}

	_, err = testQueries.PostCall(ctx, arg2)
	require.Error(t, err)
}

// TestGetCalls проверяет получение данных из stats
func TestGetCalls(t *testing.T) {
	service1 := createRandomService(t)
	service2 := createRandomService(t)
	serviceID1 := service1.ID
	serviceID2 := service2.ID

	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	userID1 := user1.ID
	userID2 := user2.ID

	ctx := context.Background()

	// Вставляем тестовые данные
	_, err := testQueries.PostCall(ctx, sqlc.PostCallParams{
		UserID:    userID1,
		ServiceID: serviceID1,
	})
	require.NoError(t, err)

	_, err = testQueries.PostCall(ctx, sqlc.PostCallParams{
		UserID:    userID2,
		ServiceID: serviceID2,
	})
	require.NoError(t, err)

	_, err = testQueries.PostCall(ctx, sqlc.PostCallParams{
		UserID:    userID1,
		ServiceID: serviceID2,
	})
	require.NoError(t, err)

	testCases := []struct {
		name   string
		params sqlc.GetStatsParams
		expect int
	}{
		{
			name: "Filter by UserID",
			params: sqlc.GetStatsParams{
				UserID:    userID1,
				ServiceID: 0, // Игнорируем ServiceID
				Limit:     10,
				Offset:    0,
			},
			expect: 2,
		},
		{
			name: "Filter by ServiceID",
			params: sqlc.GetStatsParams{
				UserID:    0, // Игнорируем UserID
				ServiceID: serviceID2,
				Limit:     10,
				Offset:    0,
			},
			expect: 2,
		},
		{
			name: "Filter by UserID and ServiceID",
			params: sqlc.GetStatsParams{
				UserID:    userID1,
				ServiceID: serviceID2,
				Limit:     10,
				Offset:    0,
			},
			expect: 1,
		},
		{
			name: "No filters (all data)",
			params: sqlc.GetStatsParams{
				UserID:    0, // Игнорируем UserID
				ServiceID: 0, // Игнорируем ServiceID
				Limit:     10,
				Offset:    0,
			},
			expect: 10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stats, err := testQueries.GetStats(ctx, tc.params)
			require.NoError(t, err)
			require.Len(t, stats, tc.expect)
		})
	}
}
