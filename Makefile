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
	oapi-codegen -config server.cfg.yaml openapi/reference/AssetService.yaml > internal/api/server.gen.go
	oapi-codegen -config servertypes.cfg.yaml openapi/reference/AssetService.yaml > internal/api/servertypes.gen.go

migrate:
	goose postgres $(POSTGRES_CONN) up -dir database/migrations

FORCE: ;
