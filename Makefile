BINARY_NAME := schema-first-price-history

build: FORCE
	go build -o build/$(BINARY_NAME) server/*

run:
	build/$(BINARY_NAME)

generate:
	sqlc generate
	go run github.com/jmattheis/goverter/cmd/goverter@v1.0.0 gen internal/...

migrate:
	goose $(POSTGRES_CONN) up -dir migrations/ postgres "user=postgres password=postgres dbname=postgres sslmode=disable"

FORCE: ;
