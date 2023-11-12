package asset

import (
	"context"

	"github.com/NinjaPerson24119/schemafirstpricehistory/internal/asset"
	"github.com/NinjaPerson24119/schemafirstpricehistory/internal/asset/generated"
	assetrepository "github.com/NinjaPerson24119/schemafirstpricehistory/internal/asset/repository"
	"github.com/gofrs/uuid"
)

type AssetService struct {
	assetRepository assetrepository.IAssetRepository
}

func NewAssetService(assetRepository assetrepository.IAssetRepository) IAssetService {
	return &AssetService{
		assetRepository: assetRepository,
	}
}

func (s *AssetService) ListAssets(ctx context.Context) ([]asset.Asset, error) {
	dbModels, err := s.assetRepository.ListAssets(ctx)
	if err != nil {
		return nil, err
	}

	converter := generated.ConverterImpl{}
	return converter.FromRepoModels(dbModels), nil
}

func (s *AssetService) GetAsset(ctx context.Context, assetID string) (asset.Asset, error) {
	u, err := uuid.FromString(assetID)
	if err != nil {
		return asset.Asset{}, err
	}

	dbModel, err := s.assetRepository.GetAsset(ctx, u)
	if err != nil {
		return asset.Asset{}, err
	}

	converter := generated.ConverterImpl{}
	return converter.FromRepoModel(dbModel), nil
}

func (s *AssetService) CreateAsset(ctx context.Context, model asset.Asset) (asset.Asset, error) {
	converter := generated.ConverterImpl{}

	createParams := converter.ToCreateParams(model)
	dbModel, err := s.assetRepository.CreateAsset(ctx, createParams)
	if err != nil {
		return asset.Asset{}, err
	}
	return converter.FromRepoModel(dbModel), nil
}

func (s *AssetService) UpdateAsset(ctx context.Context, model asset.Asset) (asset.Asset, error) {
	converter := generated.ConverterImpl{}

	updateParams, err := converter.ToUpdateParams(model)
	if err != nil {
		return asset.Asset{}, nil
	}

	dbModel, err := s.assetRepository.UpdateAsset(ctx, updateParams)
	if err != nil {
		return asset.Asset{}, err
	}
	return converter.FromRepoModel(dbModel), nil
}

func (s *AssetService) DeleteAsset(ctx context.Context, assetID string) error {
	u, err := uuid.FromString(assetID)
	if err != nil {
		return err
	}
	return s.assetRepository.DeleteAsset(ctx, u)
}
