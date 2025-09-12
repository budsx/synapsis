curl -X POST http://localhost:8003/v1/order/create-order \
-H "Content-Type: application/json" \
-H "x-idempotency-key: 1234567890" \
-d '{"product_id": 1, "quantity": 1}'