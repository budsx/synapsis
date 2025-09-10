# Inventory Service

A microservice for managing inventory with both gRPC and REST API endpoints.

## Features

- **gRPC Server**: High-performance binary protocol for internal service communication
- **REST API**: HTTP/JSON API for external clients via gRPC-Gateway
- **PostgreSQL**: Database for persistent storage
- **Docker Support**: Containerized deployment

## API Endpoints

### gRPC Endpoints (Port 8000)
- `CheckStock(product_id)` - Check available stock for a product
- `ReserveStock(product_id, quantity)` - Reserve stock for a product
- `ReleaseStock(product_id, quantity)` - Release reserved stock
- `GetProductByID(product_id)` - Get product details by ID

### REST Endpoints (Port 8001)
- `GET /v1/products/{product_id}` - Get product details
- `POST /v1/products/{product_id}/check-stock` - Check stock
- `POST /v1/products/{product_id}/reserve` - Reserve stock
- `POST /v1/products/{product_id}/release` - Release stock

## Prerequisites

- Go 1.23+
- Docker and Docker Compose
- PostgreSQL (or use Docker)

## Running the Application

### 1. Start the Database
```bash
docker-compose up -d postgres
```

### 2. Run Database Migrations
```bash
# Connect to the database and run the migration
psql -h localhost -p 5432 -U inventory_user -d inventory_db -f migration/001_create_products_table.sql
```

### 3. Start the Service
```bash
go run main.go
```

The service will start both:
- gRPC server on port 8000
- REST API server on port 8001

## Environment Variables

Create a `.env` file with the following variables:

```env
SERVICE_NAME=inventory-service
DB_DRIVER=postgres
DB_HOST=postgres
DB_PORT=5432
DB_USER=inventory_user
DB_PASSWORD=inventory_password
DB_NAME=inventory_db
LOG_LEVEL=-1
GRPC_PORT=8000
REST_PORT=8001
```

## Testing the API

### REST API Example
```bash
# Check stock
curl -X POST http://localhost:8001/v1/products/1/check-stock

# Get product details
curl http://localhost:8001/v1/products/1
```

### gRPC Example
Use a gRPC client like `grpcurl` or `evans` to test the gRPC endpoints:

```bash
# Install grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# List services
grpcurl -plaintext localhost:8000 list

# Call CheckStock
grpcurl -plaintext -d '{"product_id": 1}' localhost:8000 inventory.InventoryService/CheckStock
```

## Project Structure

```
├── config/          # Configuration management
├── entity/          # Domain entities
├── handler/         # gRPC handlers
├── migration/       # Database migrations
├── proto/           # Protocol buffer definitions
├── repository/      # Data access layer
├── services/        # Business logic
├── transport/       # Transport layer (gRPC/REST)
└── main.go         # Application entry point
```

## Development

### Generate Protocol Buffers
```bash
cd proto
make generate
```

### Build
```bash
go build -o inventory-service main.go
```

### Run Tests
```bash
go test ./...
```