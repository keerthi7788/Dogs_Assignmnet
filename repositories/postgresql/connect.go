package postgresql

import (
	"context"
	"fmt"
	"time"

	// Local packages
	configs "dogs-service/config"

	// External packages
	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect establishes a connection to the PostgreSQL database
func Connect(ctx context.Context, dbConfig configs.Postgres) (*pgxpool.Pool, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname, dbConfig.SslMode)

	// Open a new database connection
	conn, err := pgxpool.New(ctx, psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New error: %v", err)
	}

	// ping
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := conn.Ping(pingCtx); err != nil {
		return nil, fmt.Errorf("unable to connect to postgres: %v", err)
	}

	return conn, nil
}
