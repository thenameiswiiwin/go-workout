# 🏋️‍♂️ Go Workout API

A RESTful API built in Go for managing workout routines, user accounts, and authentication using Postgres. Includes token-based authentication, workout entry tracking, and a modular, testable code structure.

---

## 📁 Project Structure

```
go-workout/
├── internal/
│   ├── api/                # HTTP handlers for routes
│   ├── app/                # App configuration
│   ├── middleware/         # Authentication middleware
│   ├── routes/             # Route definitions
│   ├── store/              # Postgres DB logic
│   ├── tokens/             # Token generation and validation
│   └── utils/              # Utility helpers
├── migrations/             # Goose DB migration files
├── docker-compose.yml      # DB container setup
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## 🚀 Features

- 🔐 User Registration & Login
- 🔑 Secure Token-based Authentication
- 📋 Create, Read, Update Workouts & Entries
- 🧪 Unit Tested Workout Store
- 🐘 PostgreSQL DB with Goose Migrations
- 🧱 Middleware for protected routes
- 🐳 Docker support for development

---

## 🛠️ Tech Stack

- Go 1.21+
- PostgreSQL
- Goose (DB migration tool)
- Docker & Docker Compose

---

## 🧪 Setup

### 1. Clone the repo

```
git clone https://github.com/thenameiswiiwin/go-workout.git
cd go-workout
```

### 2. Start Postgres using Docker

```
docker-compose up --build
```

### 3. Apply Migrations

Install goose if not already:

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Then run:

```
goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/go_workout?sslmode=disable" up
```

### 4. Run the Server

```
go run main.go
```

---

## 🧪 Running Tests

```
go test ./internal/store
```

---

## 🧪 Example API Requests

### 🔐 Register a User

```
curl -X POST "http://localhost:8080/users" \
     -H "Content-Type: application/json" \
     -d '{
          "username": "thenameiswiiwin",
          "email": "thenameiswiiwin@example.com",
          "password": "SecureP@ssword123",
          "bio": "Fitness enthusiast and software developer"
        }'
```

### 🔑 Login for Token

```
curl -X POST "http://localhost:8080/tokens/authentication" \
     -H "Content-Type: application/json" \
     -d '{
          "username": "thenameiswiiwin",
          "password": "SecureP@ssword123"
        }'
```

### 🏋️‍♂️ Create Workout

```
curl -X POST "http://localhost:8080/workouts" \
     -H "Authorization: Bearer <your-token>" \
     -H "Content-Type: application/json" \
     -d '{
          "title": "Morning Cardio",
          "description": "A light 30-minute jog to start the day.",
          "duration_minutes": 30,
          "calories_burned": 300,
          "entries": [
              {
                  "exercise_name": "Jogging",
                  "sets": 1,
                  "duration_seconds": 1800,
                  "weight": 0,
                  "notes": "Maintain a steady pace",
                  "order_index": 1
              }
          ]
        }'
```

### ✏️ Update Workout

```
curl -X PUT "http://localhost:8080/workouts/6" \
     -H "Authorization: Bearer <your-token>" \
     -H "Content-Type: application/json" \
     -d '{
          "title": "Updated Cardio",
          "description": "A relaxed 45-minute walk after dinner.",
          "duration_minutes": 45,
          "calories_burned": 250,
          "entries": [
              {
                  "exercise_name": "Walking",
                  "sets": 1,
                  "duration_seconds": 2700,
                  "weight": 0,
                  "notes": "Keep a steady pace",
                  "order_index": 1
              }
          ]
        }'
```
