package service

import (
	"context"

	"github.com/Iemaduddin/goweb/backend-go/internal/model"
	"github.com/Iemaduddin/goweb/backend-go/internal/repository"
)

type AssetService interface {
	CreateAsset(ctx context.Context, asset *model.Asset) error
	GetAssetByID(ctx context.Context, id int64) (*model.Asset, error)
	FindAllAssets(ctx context.Context) ([]model.Asset, error)
	UpdateAsset(ctx context.Context, asset *model.Asset) error
	DeleteAsset(ctx context.Context, id int64) error
	ActivateAsset(ctx context.Context, id int64, isActive bool) error
}

type assetService struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{
		assetRepo: repo,
	}
}

func (s *assetService) CreateAsset(ctx context.Context, asset *model.Asset) error {
	return s.assetRepo.CreateAsset(ctx, asset)
}

func (s *assetService) GetAssetByID(ctx context.Context, id int64) (*model.Asset, error) {
	return s.assetRepo.GetAssetByID(ctx, id)
}

func (s *assetService) FindAllAssets(ctx context.Context) ([]model.Asset, error) {
	return s.assetRepo.FindAllAssets(ctx)
}

func (s *assetService) UpdateAsset(ctx context.Context, asset *model.Asset) error {
	return s.assetRepo.UpdateAsset(ctx, asset)
}

func (s *assetService) DeleteAsset(ctx context.Context, id int64) error {
	return s.assetRepo.DeleteAsset(ctx, id)
}

func (s *assetService) ActivateAsset(ctx context.Context, id int64, isActive bool) error {
	return s.assetRepo.ActivateAsset(ctx, id, isActive)
}
