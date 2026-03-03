package postgresql

import (
	"context"
	"dogs-service/models"
	"fmt"
)

// GetAllDogs returns all dogs from the database
func (r *DogRepository) GetAllDogs(ctx context.Context) ([]models.Dog, error) {
	query := `SELECT id, breed, sub_breed FROM dogs ORDER BY id`

	rows, err := r.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dogs []models.Dog
	for rows.Next() {
		var dog models.Dog
		if err := rows.Scan(&dog.ID, &dog.Breed, &dog.SubBreed); err != nil {
			return nil, err
		}
		dogs = append(dogs, dog)
	}

	return dogs, nil
}

// GetByID returns a dog by its ID
func (r *DogRepository) GetByID(ctx context.Context, id int) (models.Dog, error) {
	var dog models.Dog
	query := `SELECT id, breed, sub_breed FROM dogs WHERE id=$1`
	err := r.conn.QueryRow(ctx, query, id).Scan(&dog.ID, &dog.Breed, &dog.SubBreed)
	return dog, err
}

// Update updates an existing dog's details
func (r *DogRepository) Update(ctx context.Context, dog models.Dog) error {
	query := `UPDATE dogs SET breed=$1, sub_breed=$2 WHERE id=$3`
	_, err := r.conn.Exec(ctx, query, dog.Breed, dog.SubBreed, dog.ID)
	return err
}

// Delete removes a dog by ID
func (r *DogRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM dogs WHERE id=$1`
	result, err := r.conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected() // single return, no error

	if rowsAffected == 0 {
		return fmt.Errorf("dog not found")
	}

	return nil
}
