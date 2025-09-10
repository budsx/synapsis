compose-up:
	docker compose up --build

compose-down:
	docker compose down -v

compose-service:
	docker compose up --build inventory-service -d

gen_mock:
	cd repository/interfaces && mockgen -source=interface.go -package=interfaces -destination=interface_mock.go

test:
	go test ./services/... -coverprofile=coverage.out

test-coverage:
	go test ./services/... -coverprofile=coverage.out
	go tool cover -html=coverage.out
