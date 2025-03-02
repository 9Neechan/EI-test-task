package app

import (
	"log"

	config "github.com/9Neechan/EI-test-task/api-gateway/internal/config"
	gclient "github.com/9Neechan/EI-test-task/api-gateway/internal/grpc_client"
	hserver "github.com/9Neechan/EI-test-task/api-gateway/internal/http_api"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// TODO delete not needed fields
type serviceProvider struct {
	grpcConfig  config.GRPCConfig
	httpConfig  config.HTTPConfig
	gconn       *grpc.ClientConn
	gclientImpl *gclient.GRPCClient
	router      *gin.Engine
	hserver     *hserver.Server
}

func newServiceProvider() *serviceProvider {
	sp := &serviceProvider{}
	sp.HTTPConfig()
	sp.GRPCConfig()

	return sp
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// CloseGRPC закрывает gRPC-соединение
/*func (s *serviceProvider) CloseGRPC() {
	if s.gconn != nil {
		s.gconn.Close()
		fmt.Println("✅ gRPC-соединение закрыто")
	}
}*/
