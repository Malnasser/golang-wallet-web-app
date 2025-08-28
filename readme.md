# Go Gin + GORM Project

A backend boilerplate project built with [Go](https://go.dev/), [Gin](https://gin-gonic.com/), [GORM](https://gorm.io/), [golang-migrate](https://github.com/golang-migrate/migrate), and containerized with [Docker](https://www.docker.com/) + [docker-compose](https://docs.docker.com/compose/).  
Configuration is handled via a `.env` file for flexibility.

---

## ⚙️ Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- [Docker](https://www.docker.com/get-started) & [docker-compose](https://docs.docker.com/compose/)
- [golang-migrate](https://github.com/golang-migrate/migrate) (optional, for local DB migrations)

---

## 🚀 Setup & Run

### 1. Clone repo

```bash
git clone https://github.com/yourusername/yourproject.git
cd yourproject
```

### 2. Configure .env

Create a .env file in the project root, check .env.example for referenence

.env local development:

```
# Server Configuration
GIN_MODE=release
PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=app
DB_SSLMODE=disable

# JWT Secret (for authentication)
JWT_SECRET=your-super-secret-jwt-key-here

# Other environment variables
APP_ENV=development
```

### 3. Run with Docker

```bash
docker-compose up --build
```

---

## 🗄️ Database & Migrations

Migrations are managed with golang-migrate.

### Run migrations inside Docker

```bash
docker-compose run --rm app migrate -path ./migrations -database "postgres://postgres:postgres@db:5432/mydb?sslmode=disable" up
```

### Rollback last migration

```bash
docker-compose run --rm app migrate -path ./migrations -database "postgres://postgres:postgres@db:5432/mydb?sslmode=disable" down 1
```
