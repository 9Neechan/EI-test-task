package api

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	limiterMemory "github.com/ulule/limiter/v3/drivers/store/memory"

	gclient "github.com/9Neechan/EI-test-task/api-gateway/internal/grpc_client"
)

// Server represents the HTTP server
type Server struct {
	adress  string
	router  *gin.Engine
	gClient *gclient.GRPCClient
}

// NewServer creates a new instance of the Server
func NewServer(adress string, gClient *gclient.GRPCClient) *Server {
	server := &Server{
		adress:  adress,
		gClient: gClient,
	}

	server.setupRouter()
	return server
}

// setupRouter sets up the Gin router with rate limiting middleware
func (server *Server) setupRouter() {
	router := gin.Default()

	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  100,
	}
	store := limiterMemory.NewStore()
	limiterMiddleware := mgin.NewMiddleware(limiter.New(store, rate))

	limitedRoutes := router.Group("/").Use(limiterMiddleware)
	limitedRoutes.POST("/call", server.postCall)
	limitedRoutes.GET("/calls", server.getStats)
	limitedRoutes.POST("/service", server.createService)

	server.router = router
}

// Start starts the server
func (server *Server) Start() error {
	log.Println("on", server.adress)
	return server.router.Run(server.adress)
}

// errorResponse returns a Gin response with an error message
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
