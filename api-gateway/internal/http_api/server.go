package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	limiterMemory "github.com/ulule/limiter/v3/drivers/store/memory"
)

type Server struct {
	adress string
	router *gin.Engine
}

func NewServer(adress string) (*Server) {
	server := &Server{
		adress: adress,
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  100,
	}
	store := limiterMemory.NewStore()
	limiterMiddleware := mgin.NewMiddleware(limiter.New(store, rate))

	limitedRoutes := router.Group("/").Use(limiterMiddleware)
	limitedRoutes.POST("/call", server.postCall)         // http://localhost:8080/call?user_id=123&service_id=456
	limitedRoutes.GET("/calls", server.getStats)         // http://localhost:8080/calls?user_id=123&service_id=456&limit=10&offset=20
	limitedRoutes.POST("/service", server.createService) // http://localhost:8080/service   POST /

	server.router = router
}

func (server *Server) Start() error {
	log.Println("on", server.adress)
	return server.router.Run(server.adress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
