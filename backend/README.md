# Interview Clone - Backend (Golang + PostgreSQL)

## Prerequisites
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL (or use Docker)

## Setup

### 1. Start PostgreSQL Database
```bash
cd backend
docker-compose up -d
```

### 2. Configure Environment
```bash
cp .env.example .env
```

Edit `.env` with your database credentials.

### 3. Install Dependencies
```bash
go mod download
go mod tidy
```

### 4. Run the Server
```bash
go run main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

### Public Routes
- `GET /health` - Health check
- `GET /api/categories` - Get all categories
- `GET /api/questions` - Get questions (with filters: category, difficulty, limit, offset)
- `GET /api/questions/:id` - Get single question with answers
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login user

### Protected Routes (require Bearer token)
- `POST /api/progress` - Save user progress
- `GET /api/progress` - Get user progress
- `PUT /api/user/profile` - Update user profile

## Database Schema

### Tables
- `users` - User accounts
- `categories` - Question categories (HTML, CSS, JavaScript, etc.)
- `questions` - Interview questions
- `answers` - Multiple choice answers
- `user_progress` - Track user learning progress
- `bookmarks` - Saved questions

## Project Structure
```
backend/
├── main.go              # Application entry point
├── config/
│   ├── config.go        # Configuration loading
│   └── migrations.go    # Database migrations
├── handlers/
│   └── handlers.go      # HTTP request handlers
├── middleware/
│   └── auth.go          # Authentication middleware
├── models/
│   └── models.go        # Data models
├── go.mod               # Go module definition
├── docker-compose.yml   # Docker setup for PostgreSQL
└── .env.example         # Environment variables template
```

## Development

### Run Tests
```bash
go test ./...
```

### Build Binary
```bash
go build -o interview-clone
```

## Notes
- Password hashing and JWT validation are placeholders - implement proper security in production
- CORS is configured to allow all origins - restrict in production
- Add proper error handling and logging as needed
