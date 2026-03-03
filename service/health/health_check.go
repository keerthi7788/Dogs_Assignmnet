package health

import (
	// Go internal packages
	"context"

	// External packages
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type HealthCheckerService struct {
	logger                             *zap.Logger
	postgresClient, postgresTestClient *pgxpool.Pool
}

// NewService creates a new HealthCheckerService instance and returns the instance.
func NewService(logger *zap.Logger, postgresClient, postgresTestClient *pgxpool.Pool) *HealthCheckerService {
	return &HealthCheckerService{
		logger:             logger,
		postgresClient:     postgresClient,
		postgresTestClient: postgresTestClient,
	}
}

// Health checks the health of the database connections and returns true if all the connections are healthy.
func (h *HealthCheckerService) Health(ctx context.Context) bool {

	// Ping the database to verify the connection
	if postgresErr := h.postgresClient.Ping(ctx); postgresErr != nil {
		h.logger.Error("Postgres ping failed", zap.Error(postgresErr))
		return false
	}

	if postgresTestErr := h.postgresTestClient.Ping(ctx); postgresTestErr != nil {
		h.logger.Error("Postgres Test ping failed", zap.Error(postgresTestErr))
		return false
	}

	return true
}
