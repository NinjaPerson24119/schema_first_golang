BINARY_NAME := schema-first-price-history

build: FORCE
	go build -o build/$(BINARY_NAME) server/*

run:
	build/$(BINARY_NAME)

generate:
	sqlc generate
	goverter gen \
		-g "extend github.com/NinjaPerson24119/schemafirstpricehistory/internal/typeconversions:.*" \
		github.com/NinjaPerson24119/schemafirstpricehistory/internal/asset

migrate:
	goose postgres $(POSTGRES_CONN) up -dir database/migrations

FORCE: ;
