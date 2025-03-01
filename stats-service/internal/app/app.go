package app

import (
	"context"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	desc "github.com/9Neechan/EI-test-task/api/pb"
	"github.com/9Neechan/EI-test-task/stats-service/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
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
	return a.runGRPCServer()
}

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

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("../../configs/cfg.env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider() // Передаем конфиг
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	// Создаем gRPC сервер
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	// Регистрируем рефлексию для gRPC UI (например, grpcurl)
	reflection.Register(a.grpcServer)

	service := a.serviceProvider.GapiImpl()
	desc.RegisterStatsServiceServer(a.grpcServer, service)

	return nil
}

func (a *App) runGRPCServer() error {
	fmt.Println(a.serviceProvider) //!!!!!
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
