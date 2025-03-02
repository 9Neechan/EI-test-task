package api

import (
	"context"
	"net/http"
	"time"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type getStatsHttpRequest struct {
	UserID    int64 `form:"user_id"`
	ServiceID int64 `form:"service_id"`
	Limit     int32 `form:"limit"`
	Offset    int32 `form:"offset"`
}

func (server *Server) getStats(ctx *gin.Context) {
	var req getStatsHttpRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Формируем gRPC-запрос
	grpcReq := &desc.GetStatsRequest{
		UserId:    &req.UserID,
		ServiceId: &req.ServiceID,
		Limit:     req.Limit,
		Page:      req.Offset,
	}

	// Устанавливаем контекст с таймаутом
	gCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Вызываем gRPC-клиент
	resp, err := server.gClient.GetStats(gCtx, grpcReq)
	if err != nil {
		grpcStatus, _ := status.FromError(err)
		ctx.JSON(mapGRPCToHTTPStatus(grpcStatus.Code()), errorResponse(err))
		return
	}

	// Отправляем успешный JSON-ответ с полученной статистикой
	ctx.JSON(http.StatusOK, resp)
}
