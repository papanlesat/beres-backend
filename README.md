## Overview  
This API server is built with the Gin web framework for high-performance HTTP routing, uses GORM as its ORM layer with MySQL, manages configuration via Viper for 12-Factor compatibility, logs through Logrus for structured output, and provides Docker Compose and Makefile support for easy local development and deployment. It is based on the [gin-boilerplate template by akmamun](https://github.com/akmamun/gin-boilerplate), licensed under Apache-2.0.  

## Features  
- **HTTP Server & Routing**: Powered by Gin, with middleware support, route grouping, and blazing performance.  
- **ORM Layer**: GORM offers developer-friendly ORM abstractions, associations, and migrations.  
- **Configuration**: Viper loads settings from `.env` files or environment variables, supporting JSON/YAML/TOML and 12-Factor practices.  
- **Logging**: Logrus provides leveled, structured logging consistent with standard library API.  
- **Containerization**: Docker Compose setup for local MySQL service and live-reload development workflow.  
- **Build Automation**: A Makefile automates building, testing, and other tasks to streamline your workflow.  

## Prerequisites  
- Go **1.23+** (modules enabled)  
- Docker & Docker Compose (for local development)
- MySQL 8.0 or compatible (for production or direct installs)  

## Getting Started  

### 1. Clone & Configure  
```bash
git clone https://github.com/papanlesat/beres-backend.git your-project
cd your-project
cp .env.example .env
# Edit .env to set DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, SERVER_HOST, SERVER_PORT, DEBUG
```

### 2. Local Development with Docker  
```bash
docker-compose -f docker-compose-dev.yml up --build
# - MySQL runs on DB_HOST=db, port 3306
# - App runs on SERVER_HOST:SERVER_PORT per .env
```

### 3. Build & Run Manually  
```bash
# Install dependencies
go mod tidy
# Run migrations and server
go run main.go
```

## Configuration  
| Env Variable      | Description                           | Default  |
|-------------------|---------------------------------------|----------|
| DB_HOST           | MySQL hostname                        | localhost|
| DB_PORT           | MySQL port                            | 3306     |
| DB_USER           | MySQL username                        | root     |
| DB_PASSWORD       | MySQL password                        | (none)   |
| DB_NAME           | Database name                         | app      |
| SERVER_HOST       | Bind address                          | 0.0.0.0  |
| SERVER_PORT       | HTTP port                             | 8000     |
| DEBUG             | Gin debug mode (true/false)           | false    |

## Project Structure  
```
├── config
│   ├── config.go       # Viper loader
│   └── db.go           # MySQL DSN builder
├── controllers        # HTTP handlers
├── infra
│   ├── database       # GORM init (MySQL only)
│   └── logger         # Logrus setup
├── migrations         # AutoMigrate models
├── models             # GORM models
├── repository         # Generic CRUD wrappers
├── routers            # Route definitions & middleware
├── helpers            # Response structs, token utils
├── docker-compose-*.yml
├── Dockerfile*        # Container builds
├── Makefile           # build & dev commands
└── main.go            # entrypoint
```

## Usage Examples  
- **List Sections**  
  ```bash
  curl http://localhost:8000/sections
  ```  
- **Auth: Register**  
  ```bash
  curl -X POST http://localhost:8000/register \
    -H "Content-Type: application/json" \
    -d '{"name":"Jane","email":"jane@ex.com","password":"secret"}'
  ```  
- **Auth: Login & Token**  
  ```bash
  curl -X POST http://localhost:8000/login \
    -H "Content-Type: application/json" \
    -d '{"email":"jane@ex.com","password":"secret","token_name":"app"}'
  ```  
  Returns `{ "token": "<raw_token>" }` for use in `Authorization: Bearer <raw_token>`.  

## Contributing  
This project extends the [gin-boilerplate by akmamun](https://github.com/akmamun/gin-boilerplate). Feel free to submit issues and pull requests following the original template’s guidelines.

## License  
Apache License 2.0. See [LICENSE](LICENSE).
