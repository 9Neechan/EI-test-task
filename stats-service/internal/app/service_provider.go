package app

import (
	"database/sql"
	"log"

	config "github.com/9Neechan/EI-test-task/stats-service/internal/config"
	repo "github.com/9Neechan/EI-test-task/stats-service/internal/db/repository"
	sqlc "github.com/9Neechan/EI-test-task/stats-service/internal/db/sqlc"
	gapi "github.com/9Neechan/EI-test-task/stats-service/internal/gapi"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig
	dbConfig   config.DBConfig
	db         *sql.DB
	repository sqlc.Querier         //dao
	gapiImpl   *gapi.Implementation // api
	//userService service.UserService // service
}

func newServiceProvider() *serviceProvider {
	sp := &serviceProvider{}
	sp.DBConfig()
	sp.GRPCConfig()

	return sp
}

func (s *serviceProvider) DBConfig() config.DBConfig {
	if s.dbConfig == nil {
		cfg, err := config.NewDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		s.dbConfig = cfg
	}

	return s.dbConfig
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

func (s *serviceProvider) Repository() sqlc.Querier {
	// Открываем соединение с БД
	db, err := sql.Open(s.dbConfig.Driver(), s.dbConfig.Source())
	if err != nil {
		log.Fatalf("❌ Failed to open database: %v", err)
		return nil
	}

	// Проверяем соединение с БД
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
		return nil
	}

	log.Println("✅ Database connected successfully!")
	s.db = db

	if s.repository == nil {
		s.repository = repo.NewSQLStore(s.db)
	}
	return s.repository
}

func (s *serviceProvider) GapiImpl() *gapi.Implementation {
	if s.gapiImpl == nil {
		s.gapiImpl = gapi.NewImplementation(s.Repository())
	}

	return s.gapiImpl
}
