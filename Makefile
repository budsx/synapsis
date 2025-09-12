compose-up:
	docker compose up --build

compose-down:
	docker compose down -v

compose-service-inventory:
	docker compose up --build inventory-service -d

compose-service-order:
	docker compose up --build order-service -d

compose-service-all:
	docker compose up --build inventory-service order-service -d
