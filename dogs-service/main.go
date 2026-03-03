package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"dogs-service/config"
	"dogs-service/http"
	"dogs-service/http/handlers"
	"dogs-service/repositories/postgresql"
	"dogs-service/seeds"
	"dogs-service/service"

	_ "github.com/lib/pq"
)

func main() {

	log.Println("Starting Dogs Service...")

	cfg := config.Load()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	// Step 1: Connect to PostgreSQL
	db, err := Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Step 2: Load dogs.json and insert into PostgreSQL
	err = seeds.SeedDogs(db, "dogs.json")
	if err != nil {
		log.Fatal("Seed failed:", err)
	}

	log.Println("Bulk data inserted successfully")

	// Step 3: Start application normally
	repo := postgresql.NewDogRepository(db)
	svc := service.NewDogService(repo)
	handler := handlers.NewDogHandler(svc)

	server := http.NewServer(handler)

	log.Println("Server listening on port:", cfg.Port)

	if err := server.Listen(ctx, cfg.Port); err != nil {
		log.Fatal(err)
	}
}

func Connect(cfg config.Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
