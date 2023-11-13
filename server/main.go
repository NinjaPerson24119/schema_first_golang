package main

import (
	"context"
	"net/http"
	"os"

	"github.com/NinjaPerson24119/schemafirstpricehistory/internal/api"
	assetrepository "github.com/NinjaPerson24119/schemafirstpricehistory/internal/asset/repository"
	assetservice "github.com/NinjaPerson24119/schemafirstpricehistory/internal/asset/service"
	"github.com/NinjaPerson24119/schemafirstpricehistory/internal/database"
)

func main() {
	ctx := context.Background()

	db, err := database.NewDatabase(ctx, os.Getenv("POSTGRES_CONN"))
	if err != nil {
		os.Exit(databaseErrorCode)
	}
	defer db.Close()

	assetRepository := assetrepository.NewAssetRepository(db)
	assetService := assetservice.NewAssetService(assetRepository)

	// TODO
	//r := chi.NewRouter()
	//r.Mount("/", Handler(&myApi))

	server := api.NewServer(assetService)
	err = http.ListenAndServe(":7000", api.Handler(server))
	if err != nil {
		os.Exit(serverErrorCode)
	}

	os.Exit(successCode)
}
