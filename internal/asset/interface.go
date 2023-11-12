package asset

import "context"

type IAssetService interface {
	ListAssets(ctx context.Context) ([]Asset, error)
	GetAsset(ctx context.Context, assetID string) (Asset, error)
	CreateAsset(ctx context.Context, asset Asset) (Asset, error)
	UpdateAsset(ctx context.Context, asset Asset) (Asset, error)
	DeleteAsset(ctx context.Context, assetID string) error
}
