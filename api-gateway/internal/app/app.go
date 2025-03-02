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

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
	gClient         *gclient.GRPCClient
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runHTTPServer()
}

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

func (a *App) initConfig(_ context.Context) error {
	//err := config.Load("../../configs/cfg.env") // local test
	err := config.Load("cfg.env") // prod
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider() // Передаем конфиг
	return nil
}

func (a *App) initGRPCClient(_ context.Context) error {
	// Подключаемся к gRPC-серверу
	conn, err := grpc.Dial(
		a.serviceProvider.grpcConfig.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("❌ Не удалось подключиться к gRPC-серверу: %w", err)
	}
	a.serviceProvider.gconn = conn

	//a.serviceProvider.gclientImpl = gclient.NewGRPCClient(conn)
	a.gClient = gclient.NewGRPCClient(conn)

	fmt.Println("✅ gRPC клиент успешно создан")

	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {
	a.httpServer = http.NewServer(a.serviceProvider.httpConfig.Address(), a.gClient)
	return nil
}

func (a *App) runHTTPServer() error {
	err := a.httpServer.Start()
	return err
}
