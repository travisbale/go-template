# Go Service Template

A production-ready template for building Go microservices with both HTTP and gRPC APIs, PostgreSQL database, and complete CI/CD setup.

---

## 🔧 Using This Template

**Delete this section after setup is complete.**

### 1. Clone and Replace Placeholders

| Placeholder | Where | Replace With |
|-------------|-------|--------------|
| `<projectname>` | Makefile, Dockerfile, README | Your service name |
| `<yourorg>` | All .go files | Your GitHub org/username |
| `app` | Go code, cmd/app/ | Your service name (optional) |

### 2. Initialize Go Module

```bash
go mod init github.com/yourorg/yourproject
go mod tidy
```

Update `sqlc.yaml` to point to your new module imports.

### 3. Define Your Schema

- Add SQL schema to `internal/db/postgres/migrations/001_init.up.sql`
- Add queries to `internal/db/postgres/queries/`
- Define Protocol Buffers in `proto/`

### 4. Generate Code

```bash
make sqlc    # Generate type-safe database code
make protoc  # Generate gRPC code
```

### 5. Build and Run

```bash
make dev
./bin/app start --database-url "postgres://postgres:password@localhost:5432/yourdb?sslmode=disable"
```

### What's Included

- **CLI framework** (`cmd/app/`) - urfave/cli for command-line interface
- **HTTP server** (`internal/api/http/`) - Chi router with middleware
- **gRPC server** (`internal/api/grpc/`) - Protocol Buffers support
- **Database layer** (`internal/db/postgres/`) - pgx connection pool with RLS support
- **Migrations** - Auto-run on startup with golang-migrate
- **Code generation** - sqlc for type-safe SQL, protoc for gRPC
- **SDK** (`sdk/`) - Public Go client library for HTTP and gRPC
- **Docker** - Multi-stage Alpine build
- **CI/CD** - GitHub Actions with linting, testing, security scanning

---

## Features

- **HTTP API**: RESTful endpoints with chi router
- **gRPC API**: Protocol Buffers service definitions
- **PostgreSQL**: Row-level security for multi-tenancy, migrations with golang-migrate
- **Type-safe database access**: Generated code with sqlc
- **Authentication**: JWT validation middleware
- **Code quality**: Comprehensive linting with golangci-lint
- **Docker**: Multi-stage builds with Alpine Linux
- **CI/CD**: GitHub Actions for testing, linting, and vulnerability scanning
- **SDK**: Public Go client library

## Quick Start

### Prerequisites

- Go 1.25+
- Docker
- PostgreSQL 16+
- Make

### Development

```bash
# Install dependencies
go mod download

# Build development binary (with debug symbols)
make dev

# Run tests
make test

# Format code
make fmt

# Lint code
make lint
```

### Database Setup

```bash
# Run PostgreSQL with Docker
docker run -d \
  --name postgres \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=<projectname> \
  -p 5432:5432 \
  postgres:16-alpine

# Migrations run automatically on startup
# Or run manually:
./bin/app migrate up --database-url "postgres://postgres:password@localhost:5432/<projectname>?sslmode=disable"
```

### Running the Service

```bash
# Build
make dev

# Start server
./bin/app start \
  --http-address ":8080" \
  --grpc-address ":9090" \
  --database-url "postgres://postgres:password@localhost:5432/<projectname>?sslmode=disable" \
  --jwt-public-key "./keys/public.pem"
```

### Code Generation

```bash
# Generate database code from SQL queries
make sqlc

# Generate gRPC/protobuf code
make protoc
```

## Project Structure

```txt
.
├── cmd/app/              # CLI entry point (urfave/cli)
│   ├── main.go          # CLI app setup
│   ├── start.go         # Server start command
│   ├── migrate.go       # Database migration commands
│   ├── version.go       # Version command
│   └── flags.go         # Shared CLI flags
├── internal/
│   ├── app/             # Server lifecycle
│   │   └── server.go    # Coordinates HTTP, gRPC, DB initialization
│   ├── api/
│   │   ├── http/        # HTTP layer (chi router)
│   │   │   └── server.go
│   │   └── grpc/        # gRPC layer
│   │       └── server.go
│   ├── db/postgres/     # Data access layer
│   │   ├── db.go        # Connection pool + tenant context helpers
│   │   ├── migrate.go   # Migration runner
│   │   ├── migrations/  # SQL schema files
│   │   ├── queries/     # SQL queries (input for sqlc)
│   │   └── sqlc/        # Generated type-safe Go code
│   └── pb/              # Generated protobuf code
├── proto/               # Protocol buffer definitions
├── sdk/                 # Public Go client library
│   ├── http_client.go   # HTTP client
│   ├── grpc_client.go   # gRPC client
│   └── types.go         # Shared types
├── Dockerfile           # Multi-stage Docker build
├── Makefile             # Build automation
└── sqlc.yaml            # Database code generation config
```

## Common Commands

```bash
make build        # Production build (optimized, stripped)
make dev          # Development build (fast, debug symbols)
make test         # Run all tests with race detector
make fmt          # Format code with gofmt/goimports
make lint         # Lint code with golangci-lint
make sqlc         # Generate database code
make protoc       # Generate gRPC code
make clean        # Clean build artifacts
make help         # Show all available commands
```

## Configuration

Environment variables (can also be passed as CLI flags):

- `HTTP_ADDRESS` - HTTP server bind address (default: `:8080`)
- `GRPC_ADDRESS` - gRPC server bind address (default: `:9090`)
- `DATABASE_URL` - PostgreSQL connection string (required)
- `JWT_PUBLIC_KEY_PATH` - Path to RSA public key PEM file (required for auth)
- `ENVIRONMENT` - Environment name (default: `development`)

## Docker

### Build

```bash
make docker-build
# Or manually:
docker build -t <projectname>:latest .
```

### Run

```bash
docker run -p 8080:8080 -p 9090:9090 \
  -e DATABASE_URL="postgres://..." \
  -e JWT_PUBLIC_KEY_PATH="/app/keys/public.pem" \
  <projectname>:latest
```

## Testing

```bash
# Run all tests
make test

# Run specific package tests
go test -v ./internal/api/http

# Run with coverage
go test -cover ./...
```

## CI/CD

The project includes a GitHub Actions workflow (`.github/workflows/ci.yml`) that runs on every push:

1. Linting (golangci-lint)
2. Build verification
3. Tests with race detector
4. Vulnerability scanning (govulncheck)
5. Code generation validation (sqlc, protobuf)
6. Docker build

## Multi-Tenancy

This template uses PostgreSQL Row-Level Security (RLS) for tenant isolation:

- Tenant ID is extracted from JWT and added to request context
- All database queries automatically filter by tenant
- Use `db.WithTenantContext(ctx, func(q *sqlc.Queries) error { ... })` for tenant-scoped operations

## Key Patterns

### Database Access

- Use `db.WithTransaction()` for non-tenant operations
- Use `db.WithTenantContext()` for tenant-scoped queries (RLS enabled)

### Adding HTTP Endpoints

Edit `internal/api/http/server.go` to add routes.

### Adding gRPC Services

1. Define in `proto/*.proto`
2. Run `make protoc`
3. Implement in `internal/api/grpc/`

## License

MIT
