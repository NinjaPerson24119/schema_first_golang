package asset

import (
	"context"

	"github.com/NinjaPerson24119/schemafirstpricehistory/internal/asset"
)

type IAssetService interface {
	ListAssets(ctx context.Context) ([]asset.Asset, error)
	GetAsset(ctx context.Context, assetID string) (asset.Asset, error)
	CreateAsset(ctx context.Context, model asset.Asset) (asset.Asset, error)
	UpdateAsset(ctx context.Context, model asset.Asset) (asset.Asset, error)
	DeleteAsset(ctx context.Context, assetID string) error
}
