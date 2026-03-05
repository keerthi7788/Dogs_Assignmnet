package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	// Local Packages
	configs "dogs-service/config"
	"dogs-service/http"
	"dogs-service/http/handlers"
	"dogs-service/logger"
	"dogs-service/repositories/postgresql"
	"dogs-service/seeds"
	"dogs-service/service"
	"dogs-service/service/health"
)

func InitializeServer(ctx context.Context, cfg configs.ApxConfig) (*http.Server, error) {
	logger.Info("Initializing PostgreSQL Connection...")

	// Postgres Connection
	postgresConn, err := postgresql.Connect(ctx, cfg.Postgres)
	if err != nil {
		logger.Error("Error connecting to PostgreSQL", err)
		return nil, err
	}
	logger.Info("PostgreSQL connection established")

	// Seed data
	if err := seeds.SeedDogs(ctx, postgresConn, "dogs.json"); err != nil {
		logger.Error("file path does not exists", err)
		return nil, err
	}
	logger.Info("Bulk dog data seeded successfully")

	// Repositories
	dogRepo := postgresql.NewDogRepository(postgresConn)

	// Services
	healthSvc := health.NewService(logger.Logger, postgresConn)
	dogSvc := service.NewDogService(dogRepo)

	// Handlers
	dogHandler := handlers.NewDogHandler(dogSvc)
	//new

	// HTTP Server
	server := http.NewServer(dogHandler, healthSvc)
	return server, nil
}

func main() {
	logger.InitLogger()
	logger.Info("Starting Dogs Service...")

	// Load config
	cfg := configs.Load()

	// Context for graceful shutdown
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// Initialize Server
	server, err := InitializeServer(ctx, cfg)
	if err != nil {
		logger.Fatal("Failed to initialize server", err)
	}

	logger.Info("Server listening on port:", cfg.Port)
	if err := server.Listen(ctx, cfg.Port); err != nil {
		logger.Fatal("Server failed to listen", err)
	}
}
