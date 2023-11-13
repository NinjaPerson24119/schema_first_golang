package api

import (
	"context"

	assetservice "github.com/NinjaPerson24119/schemafirstpricehistory/internal/asset/service"
)

type Converter interface {
}

type StrictServer struct {
	assetService assetservice.IAssetService
}

func NewServer(assetService assetservice.IAssetService) StrictServer {
	return StrictServer{
		assetService: assetService,
	}
}

func (s StrictServer) GetAssets(ctx context.Context, request GetAssetsRequestObject) (GetAssetsResponseObject, error) {
	assets, err := s.assetService.ListAssets(ctx)
	if err != nil {
		return nil, err
	}

	return GetAssets200JSONResponse{
		Assets: assets,
	}, nil
}

func (s StrictServer) PostAssetCreate(ctx context.Context, request PostAssetCreateRequestObject) (PostAssetCreateResponseObject, error)

func (s StrictServer) DeleteAssetsAssetId(ctx context.Context, request DeleteAssetsAssetIdRequestObject) (DeleteAssetsAssetIdResponseObject, error)

func (s StrictServer) GetAssetAssetId(ctx context.Context, request GetAssetAssetIdRequestObject) (GetAssetAssetIdResponseObject, error)

func (s StrictServer) PutAssetsAssetId(ctx context.Context, request PutAssetsAssetIdRequestObject) (PutAssetsAssetIdResponseObject, error)
