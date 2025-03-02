package api

import (
	"context"
	"net/http"
	"time"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type createServiceHttpRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

/*

curl -X POST "http://localhost:8080/service" \
     -H "Content-Type: application/json" \
     -d '{
           "name": "abc",
           "description": "descrpit"
         }' 
		 
*/
func (server *Server) createService(ctx *gin.Context) {
	var req createServiceHttpRequest

	// Привязываем JSON-запрос к структуре
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Формируем gRPC-запрос
	grpcReq := &desc.CreateServiceRequest{
		Name:        req.Name,
		Description: req.Description,
	}

	// Устанавливаем контекст с таймаутом
	gCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Вызываем gRPC-клиент
	resp, err := server.gClient.CreateService(gCtx, grpcReq)
	if err != nil {
		grpcStatus, _ := status.FromError(err)
		ctx.JSON(mapGRPCToHTTPStatus(grpcStatus.Code()), errorResponse(err))
		return
	}

	// Отправляем успешный JSON-ответ
	ctx.JSON(http.StatusOK, resp)
}
