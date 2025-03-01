package model

import (
	"time"
)

type Service struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
}

type CreateServiceRequest struct {
	Name        string
	Description string
}

type CreateServiceResponse struct {
	ServiceId   int64  
	Name        string
	Description string 
	CreatedAt   time.Time
}
