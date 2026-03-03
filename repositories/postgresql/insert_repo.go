package postgresql

import (
	"context"
	"dogs-service/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DogRepository struct {
	conn *pgxpool.Pool
}

func NewDogRepository(conn *pgxpool.Pool) *DogRepository {
	return &DogRepository{conn: conn}
}

// CreateDog inserts a new dog into the database
func (r *DogRepository) CreateDog(ctx context.Context, dog models.Dog) error {
	query := `
	INSERT INTO dogs (breed, sub_breed)
	VALUES ($1, $2)
	`
	_, err := r.conn.Exec(ctx, query, dog.Breed, dog.SubBreed)
	return err
}
