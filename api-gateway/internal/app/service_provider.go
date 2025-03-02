package app

import (
	"log"
	
	"google.golang.org/grpc"

	config "github.com/9Neechan/EI-test-task/api-gateway/internal/config"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig
	gconn      *grpc.ClientConn
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
