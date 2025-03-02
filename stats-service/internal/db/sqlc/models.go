// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Service struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Stat struct {
	UserID    int64 `json:"user_id"`
	ServiceID int64 `json:"service_id"`
	Count     int64 `json:"count"`
}

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
