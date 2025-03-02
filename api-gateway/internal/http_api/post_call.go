package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type postCallHttpRequest struct {
	UserID    int64 `form:"user_id" binding:"required"`
	ServiceID int64 `form:"service_id" binding:"required"`
}

func (server *Server) postCall(ctx *gin.Context) {
	var req postCallHttpRequest

	// Привязываем параметры из URI
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println("!!!", req)

	// Создаем gRPC-запрос
	grpcReq := &desc.PostCallRequest{
		UserId:    int64(req.UserID),
		ServiceId: int64(req.ServiceID),
	}

	// Устанавливаем контекст с таймаутом
	gCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Вызываем gRPC-сервис
	resp, err := server.gClient.PostCall(gCtx, grpcReq)
	if err != nil {
		grpcStatus, _ := status.FromError(err)
		ctx.JSON(mapGRPCToHTTPStatus(grpcStatus.Code()), errorResponse(err))
		return
	}

	// Отправляем успешный JSON-ответ
	ctx.JSON(http.StatusOK, gin.H{"success": resp.Success})
}
