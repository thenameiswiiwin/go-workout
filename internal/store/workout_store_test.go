package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	err = Migrate(db, "../../migrations/")
	if err != nil {
		t.Fatalf("failed to migrate database: %v", err)
	}

	_, err = db.Exec(`TRUNCATE TABLE workouts, workout_entries CASCADE`)
	if err != nil {
		t.Fatalf("failed to truncate tables: %v", err)
	}

	return db
}
