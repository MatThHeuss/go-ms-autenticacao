DB_URL=postgresql://docker:docker@localhost:5432/apigolang?sslmode=disable

start:
	go run cmd/api/main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up


.PHONY: start migrateup