package seeds

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// SeedDogs loads dogs from a JSON file and inserts them into PostgreSQL using pgxpool
func SeedDogs(ctx context.Context, db *pgxpool.Pool, filePath string) error {
	log.Println("Loading seed file:", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var dogs map[string][]string
	if err := json.NewDecoder(file).Decode(&dogs); err != nil {
		return err
	}

	query := `
	INSERT INTO dogs (breed, sub_breed)
	VALUES ($1, $2)
	ON CONFLICT DO NOTHING
	`

	for breed, subs := range dogs {
		// No sub-breeds
		if len(subs) == 0 {
			if _, err := db.Exec(ctx, query, breed, ""); err != nil {
				return err
			}
			continue
		}

		// Insert sub-breeds
		for _, sub := range subs {
			if _, err := db.Exec(ctx, query, breed, sub); err != nil {
				return err
			}
		}
	}

	log.Println("Seed completed successfully")
	return nil
}
