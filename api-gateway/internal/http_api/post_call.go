package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"

	desc "github.com/9Neechan/EI-test-task/api/pb"
)

// postCallHttpRequest represents the HTTP request structure for postCall
type postCallHttpRequest struct {
	UserID    int64 `form:"user_id" binding:"required"` 
	ServiceID int64 `form:"service_id" binding:"required"` 
}

// postCall handles the HTTP request for making a post call
func (server *Server) postCall(ctx *gin.Context) {
	var req postCallHttpRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	grpcReq := &desc.PostCallRequest{
		UserId:    int64(req.UserID),
		ServiceId: int64(req.ServiceID),
	}

	gCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := server.gClient.PostCall(gCtx, grpcReq)
	if err != nil {
		grpcStatus, _ := status.FromError(err)
		ctx.JSON(mapGRPCToHTTPStatus(grpcStatus.Code()), errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": resp.Success})
}
