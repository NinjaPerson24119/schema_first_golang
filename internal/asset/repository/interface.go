package repository

import (
	"context"

	"github.com/NinjaPerson24119/schema_first_price_history/internal/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type IAssetRepository interface {
	ListAssets(ctx context.Context) ([]sqlc.TradingAsset, error)
	GetAsset(ctx context.Context, assetID pgtype.UUID) (sqlc.TradingAsset, error)
	CreateAsset(ctx context.Context, createParams sqlc.CreateAssetParams) (sqlc.TradingAsset, error)
	UpdateAsset(ctx context.Context, updateParams sqlc.UpdateAssetParams) (sqlc.TradingAsset, error)
	DeleteAsset(ctx context.Context, assetID pgtype.UUID) error
}
