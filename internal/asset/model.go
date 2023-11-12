package asset

import (
	"time"

	"github.com/NinjaPerson24119/schemafirstpricehistory/internal/sqlc"
)

// goverter:converter
type Converter interface {
	FromRepoModels(source []sqlc.TradingAsset) []Asset
	FromRepoModel(source sqlc.TradingAsset) Asset
	ToCreateParams(source Asset) sqlc.CreateAssetParams
	ToUpdateParams(source Asset) (sqlc.UpdateAssetParams, error)
}

type Asset struct {
	AssetID     string
	Name        string
	Ticker      string
	Description string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
