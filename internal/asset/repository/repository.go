package repository

import (
	"context"

	"github.com/NinjaPerson24119/schema_first_price_history/internal/database"
	"github.com/NinjaPerson24119/schema_first_price_history/internal/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type AssetRepository struct {
	db database.IDatabase
}

func NewAssetRepository(db database.IDatabase) IAssetRepository {
	return &AssetRepository{
		db: db,
	}
}

func (s *AssetRepository) ListAssets(ctx context.Context) ([]sqlc.TradingAsset, error) {
	queries, err := s.db.GetQueries(ctx)
	if err != nil {
		return nil, err
	}
	return queries.ListAssets(ctx)
}

func (s *AssetRepository) GetAsset(ctx context.Context, assetID pgtype.UUID) (sqlc.TradingAsset, error) {
	queries, err := s.db.GetQueries(ctx)
	if err != nil {
		return sqlc.TradingAsset{}, err
	}
	return queries.GetAsset(ctx, assetID)
}

func (s *AssetRepository) CreateAsset(ctx context.Context, createParams sqlc.CreateAssetParams) (sqlc.TradingAsset, error) {
	queries, err := s.db.GetQueries(ctx)
	if err != nil {
		return sqlc.TradingAsset{}, err
	}
	return queries.CreateAsset(ctx, createParams)
}

func (s *AssetRepository) UpdateAsset(ctx context.Context, updateParams sqlc.UpdateAssetParams) (sqlc.TradingAsset, error) {
	queries, err := s.db.GetQueries(ctx)
	if err != nil {
		return sqlc.TradingAsset{}, err
	}
	return queries.UpdateAsset(ctx, updateParams)
}

func (s *AssetRepository) DeleteAsset(ctx context.Context, assetID pgtype.UUID) error {
	queries, err := s.db.GetQueries(ctx)
	if err != nil {
		return err
	}
	return queries.DeleteAsset(ctx, assetID)
}
