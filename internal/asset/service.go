package asset

import (
	"context"

	assetrepository "github.com/NinjaPerson24119/schema_first_price_history/internal/asset/repository"
)

type AssetService struct {
	assetRepository assetrepository.IAssetRepository
}

func NewAssetService(assetRepository assetrepository.IAssetRepository) IAssetService {
	return &AssetService{
		assetRepository: assetRepository,
	}
}

func (s *AssetService) ListAssets(ctx context.Context) ([]Asset, error) {
	/*
		dbModels, err := s.assetRepository.ListAssets(ctx)
		if err != nil {
			return nil, err
		}

		models := make([]Asset, len(dbModels))
		for i, v := range dbModels {
			models[i] = Asset(v)
		}
		return models, nil
	*/
	return nil, nil
}

func (s *AssetService) GetAsset(ctx context.Context, assetID string) (Asset, error) {
	return Asset{}, nil
}

func (s *AssetService) CreateAsset(ctx context.Context, asset Asset) (Asset, error) {
	return Asset{}, nil
}

func (s *AssetService) UpdateAsset(ctx context.Context, asset Asset) (Asset, error) {
	return Asset{}, nil
}

func (s *AssetService) DeleteAsset(ctx context.Context, assetID string) error {
	return nil
}
