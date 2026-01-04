# Fantasy Engine API

A production-ready REST API built with Gin framework.

## Features

- 🚀 Gin web framework
- 🗄️ PostgreSQL database with GORM
- 🔒 JWT authentication (ready to implement)
- 📝 Structured logging
- 🐳 Docker support
- 🧪 Test structure
- 📦 Clean architecture

## Project Structure

```
.
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── database/        # Database connection and migrations
│   ├── handlers/        # HTTP handlers
│   ├── middleware/      # HTTP middleware
│   ├── models/          # Data models
│   ├── router/          # Route definitions
│   ├── services/        # Business logic
│   └── utils/           # Utility functions
├── migrations/          # Database migrations (optional)
├── tests/              # Integration tests
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Docker and Docker Compose (optional)

### Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd Fantasy-Engine
```

2. Install dependencies:
```bash
go mod download
```

3. Copy environment file:
```bash
cp .env.example .env
```

4. Update `.env` with your configuration:
```env
DATABASE_URL=postgres://user:password@localhost:5432/fantasy_db?sslmode=disable
JWT_SECRET=your-secret-key-change-in-production
```

### Running Locally

#### Option 1: Using Make
```bash
make run
```

#### Option 2: Direct
```bash
go run ./cmd/server
```

#### Option 3: Using Docker Compose
```bash
make docker-up
```

The API will be available at `http://localhost:8080`

### Available Endpoints

- `GET /health` - Health check
- `GET /api/v1/ping` - Ping endpoint

## Development

### Running Tests
```bash
make test
```

### Building
```bash
make build
```

### Database Migrations
```bash
# Using golang-migrate (install separately)
make migrate-up
make migrate-down
```

## Docker

### Build Docker Image
```bash
make docker-build
```

### Run with Docker Compose
```bash
make docker-up
```

### Stop Docker Containers
```bash
make docker-down
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `ENV` | Environment (development/production) | `development` |
| `PORT` | Server port | `8080` |
| `DATABASE_URL` | PostgreSQL connection string | Required |
| `JWT_SECRET` | JWT secret key | Required in production |
| `LOG_LEVEL` | Logging level | `info` |

## Production Deployment

1. Set environment variables appropriately
2. Build the application:
```bash
make build
```

3. Run the binary:
```bash
./bin/server
```

Or use Docker:
```bash
docker build -t fantasy-engine .
docker run -p 8080:8080 --env-file .env fantasy-engine
```

## License

MIT

