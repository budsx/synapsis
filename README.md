# Synapsis Microservices

A microservices architecture for e-commerce order and inventory management built with Go, gRPC, and PostgreSQL.

![Architecture](https://i.imgur.com/rXPhvhs.png)

## System Overview and Design Choices

### Architecture
- **Microservices Coreography Pattern**
- **gRPC Communication**
- **Event-Driven**
- **Database per Service**
- **Redis for Caching**

### Services
- **Inventory Service** (Port 8000/8001): Manages product catalog and stock levels
- **Order Service** (Port 8002/8003): Handles order processing

## Setup Instructions

### Prerequisites
- Go 1.21+
- Docker and Docker Compose
- Git

### 1. Clone Repository
```bash
git clone https://github.com/budsx/synapsis.git
cd synapsis
```

### 2. Start Services
```bash
# Run Services
docker-compose up -d

# Check if all services are running
docker-compose ps
```


### 3. cURL API

#### Using cURL
```bash
# Inventory Service Tests
curl http://localhost:8001/v1/inventory/check-stock/1
curl http://localhost:8001/v1/inventory/get-product-by-id/1
curl -X POST http://localhost:8001/v1/inventory/reserve-stock \
  -H "Content-Type: application/json" \
  -d '{"product_id": "1", "quantity": "5"}'

# Order Service Tests
curl -X POST http://localhost:8003/v1/order/create-order \
  -H "Content-Type: application/json" \
  -d '{"product_id": "1", "quantity": "2", "idempotency_key": "test-001"}'
```

#### Using HTTP Files
Create `test.http` file and use VS Code REST Client extension:
```http
### Check Stock
GET http://localhost:8001/v1/inventory/check-stock/1

### Create Order
POST http://localhost:8003/v1/order/create-order
Content-Type: application/json

{
  "product_id": "1",
  "quantity": "2",
  "idempotency_key": "test-001"
}
```

### API Endpoints

#### Inventory Service (Port 8001)
- `GET /v1/inventory/check-stock/{product_id}` - Check available stock
- `GET /v1/inventory/get-product-by-id/{product_id}` - Get product details
- `POST /v1/inventory/reserve-stock` - Reserve stock
- `POST /v1/inventory/release-stock` - Release stock

#### Order Service (Port 8003)
- `POST /v1/order/create-order` - Create new order

### Stopping Services
```bash
# Stop all services
docker-compose down

# Stop and remove volumes (clean reset)
docker-compose down -v
```

