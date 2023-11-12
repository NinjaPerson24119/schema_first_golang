package repository

import (
	"context"

	"github.com/NinjaPerson24119/schemafirstpricehistory/internal/sqlc"
	"github.com/gofrs/uuid"
)

type IAssetRepository interface {
	ListAssets(ctx context.Context) ([]sqlc.TradingAsset, error)
	GetAsset(ctx context.Context, assetID uuid.UUID) (sqlc.TradingAsset, error)
	CreateAsset(ctx context.Context, createParams sqlc.CreateAssetParams) (sqlc.TradingAsset, error)
	UpdateAsset(ctx context.Context, updateParams sqlc.UpdateAssetParams) (sqlc.TradingAsset, error)
	DeleteAsset(ctx context.Context, assetID uuid.UUID) error
}
