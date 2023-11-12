package main

import (
	"context"
	"os"

	assetrepository "github.com/NinjaPerson24119/schema_first_price_history/internal/asset/repository"
	"github.com/NinjaPerson24119/schema_first_price_history/internal/database"
)

func main() {
	ctx := context.Background()

	db, err := database.NewDatabase(ctx, os.Getenv("POSTGRES_CONN"))
	if err != nil {
		os.Exit(databaseErrorCode)
	}
	defer db.Close()

	_ = assetrepository.NewAssetRepository(db)

	os.Exit(successCode)
}
