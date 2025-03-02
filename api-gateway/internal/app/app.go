package app

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/9Neechan/EI-test-task/api-gateway/internal/config"
	gclient "github.com/9Neechan/EI-test-task/api-gateway/internal/grpc_client"
	http "github.com/9Neechan/EI-test-task/api-gateway/internal/http_api"
)

// App represents the application structure
type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
	gClient         *gclient.GRPCClient
}

// NewApp initializes a new instance of the application
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run starts the HTTP server
func (a *App) Run() error {
	return a.runHTTPServer()
}

// initDeps initializes the application dependencies
func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCClient,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// initConfig loads the application configuration
func (a *App) initConfig(_ context.Context) error {
	//err := config.Load("../../configs/cfg.env") // local test
	err := config.Load("cfg.env") // prod
	if err != nil {
		return err
	}

	return nil
}

// initServiceProvider initializes the service provider
func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider() // Passes the config
	return nil
}

// initGRPCClient initializes the gRPC client
func (a *App) initGRPCClient(_ context.Context) error {
	// Connects to the gRPC server
	conn, err := grpc.Dial(
		a.serviceProvider.grpcConfig.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("Failed to connect to the gRPC server: %w", err)
	}
	a.serviceProvider.gconn = conn

	a.gClient = gclient.NewGRPCClient(conn)

	fmt.Println("gRPC client successfully created")

	return nil
}

// initHTTPServer initializes the HTTP server
func (a *App) initHTTPServer(_ context.Context) error {
	a.httpServer = http.NewServer(a.serviceProvider.httpConfig.Address(), a.gClient)
	return nil
}

// runHTTPServer starts the HTTP server
func (a *App) runHTTPServer() error {
	err := a.httpServer.Start()
	return err
}
