build:
	@go build -o bin/curiously cmd/main.go

run:
	@go run cmd/main.go

test:
	@go test -v ./...

migration: 
	@migrate create -ext sql -dir internal/model/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run internal/model/main.go up

migrate-down:
	@go run internal/model/main.go down

