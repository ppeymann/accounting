swag:
	swag init --parseDependency --parseInternal -g /server/server.go

compose:
	docker compose -f docker-compose.yaml up -d

run:
	go run cmd/accounting/main.go
