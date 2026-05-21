# Interview Clone - Luyện Phỏng Vấn IT

Clone của trang web https://luyenphongvan.online/ với stack công nghệ:
- **Backend**: Golang + Gin Framework + PostgreSQL
- **Frontend**: ReactJS + Vite + React Router

## Tính năng chính

### Backend (Golang)
- RESTful API với Gin framework
- PostgreSQL database với migrations
- Authentication với JWT
- User progress tracking
- Categories và Questions management
- CORS support

### Frontend (ReactJS)
- Responsive design
- Dark/Light theme toggle
- Multi-language support (VI/EN)
- Interactive quiz interface
- Category filtering
- Progress tracking UI
- Pricing page

## Quick Start

### 1. Start Database
```bash
cd backend
docker-compose up -d
```

### 2. Run Backend
```bash
cd backend
cp .env.example .env
go mod download
go run main.go
```

Server chạy trên `http://localhost:8080`

### 3. Run Frontend
```bash
cd frontend
npm install
npm run dev
```

App chạy trên `http://localhost:3000`

## Project Structure

```
interview-clone/
├── backend/
│   ├── main.go              # Entry point
│   ├── config/              # Configuration & migrations
│   ├── handlers/            # HTTP handlers
│   ├── middleware/          # Auth middleware
│   ├── models/              # Data models
│   ├── go.mod               # Go dependencies
│   ├── docker-compose.yml   # PostgreSQL setup
│   └── README.md
│
├── frontend/
│   ├── src/
│   │   ├── main.jsx         # React entry
│   │   ├── App.jsx          # Main component
│   │   ├── components/      # Reusable components
│   │   ├── pages/           # Page components
│   │   └── services/        # API client
│   ├── index.html
│   ├── vite.config.js
│   ├── package.json
│   └── README.md
│
└── README.md                # This file
```

## API Endpoints

### Public
- `GET /api/categories` - Get all categories
- `GET /api/questions` - Get questions (with filters)
- `GET /api/questions/:id` - Get single question
- `POST /api/auth/register` - Register user
- `POST /api/auth/login` - Login user

### Protected (requires Bearer token)
- `POST /api/progress` - Save progress
- `GET /api/progress` - Get progress
- `PUT /api/user/profile` - Update profile

## Database Schema

- **users** - User accounts
- **categories** - Question categories
- **questions** - Interview questions
- **answers** - Multiple choice answers
- **user_progress** - Learning progress
- **bookmarks** - Saved questions

## Technologies

### Backend
- Go 1.21+
- Gin Web Framework
- PostgreSQL
- lib/pq driver
- godotenv for config

### Frontend
- React 19
- Vite
- React Router DOM
- CSS Custom Properties
- Fetch API

## Development

### Backend Tests
```bash
cd backend
go test ./...
```

### Frontend Build
```bash
cd frontend
npm run build
```

## Notes

⚠️ **Security Notice**: 
- Password hashing is a placeholder - implement bcrypt in production
- JWT validation needs proper implementation
- CORS allows all origins - restrict in production
- Add rate limiting and input validation

## License

MIT License - Educational purpose only

## Credits

Inspired by https://luyenphongvan.online/

Built with ❤️ for Vietnamese developers
