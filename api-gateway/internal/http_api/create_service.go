package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"

	desc "github.com/9Neechan/EI-test-task/api/pb"
)

// createServiceHttpRequest represents the HTTP request for creating a service
type createServiceHttpRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// createService handles the HTTP request for creating a service
func (server *Server) createService(ctx *gin.Context) {
	var req createServiceHttpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	grpcReq := &desc.CreateServiceRequest{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	gCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := server.gClient.CreateService(gCtx, grpcReq)
	if err != nil {
		grpcStatus, _ := status.FromError(err)
		ctx.JSON(mapGRPCToHTTPStatus(grpcStatus.Code()), errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
