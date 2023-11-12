package asset

import (
	"time"

	"github.com/NinjaPerson24119/schema_first_price_history/internal/sqlc"
)

// goverter:converter
type Converter interface {
	ConvertItems(source []sqlc.TradingAsset) []Asset
	Convert(source sqlc.TradingAsset) Asset
}

type Asset struct {
	AssetID     string
	Name        string
	Ticker      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Description string
}
