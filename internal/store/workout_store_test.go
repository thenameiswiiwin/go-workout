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

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	store := NewPostgresWorkoutStore(db)

	tests := []struct {
		name    string
		workout *Workout
		wantErr bool
	}{
		{
			name: "valid workout",
			workout: &Workout{
				Title:           "Push Day",
				Description:     "Workout focused on pushing exercises",
				DurationMinutes: 60,
				CaloriesBurned:  500,
				Entries: []WorkoutEntry{
					{
						ExerciseName: "Bench Press",
						Sets:         3,
						Reps:         IntPtr(10),
						Weight:       FloatPtr(135.5),
						Notes:        "Felt strong",
						OrderIndex:   1,
					},
				},
			},
			wantErr: false,
		},
	}
}

func IntPtr(i int) *int { return &i }

func FloatPtr(f float64) *float64 { return &f }
