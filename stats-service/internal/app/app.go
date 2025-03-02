package app

import (
	"context"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	"github.com/9Neechan/EI-test-task/stats-service/internal/config"
)

// App is the main application struct
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp creates a new instance of the App
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run starts the application
func (a *App) Run() error {
	return a.runGRPCServer()
}

// initDeps initializes the dependencies of the application
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// initConfig initializes the configuration of the application
func (a *App) initConfig(_ context.Context) error {
	//err := config.Load("../../configs/cfg.env") // local test
	err := config.Load("cfg.env") // prod
	if err != nil {
		return err
	}

	return nil
}

// initServiceProvider initializes the service provider of the application
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider() // Pass the config
	return nil
}

// initGRPCServer initializes the gRPC server of the application
func (a *App) initGRPCServer(_ context.Context) error {
	// Create a gRPC server
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	service := a.serviceProvider.GapiImpl()
	desc.RegisterStatsServiceServer(a.grpcServer, service)

	return nil
}

// runGRPCServer starts the gRPC server of the application
func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
