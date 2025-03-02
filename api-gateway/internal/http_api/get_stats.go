package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getStatsRequest struct {
	UserID    int64 `json:"user_id"`
	ServiceID int64 `json:"service_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

// curl -X GET "http://localhost:8080/calls?user_id=123&service_id=456&page=1&limit=10"

func (server *Server) getStats(ctx *gin.Context) {
	var req postCallRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	stats := struct{}{}

	/*arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}*/

	ctx.JSON(http.StatusOK, stats)
}
