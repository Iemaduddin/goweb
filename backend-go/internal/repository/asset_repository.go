package repository

import (
	"context"
	"database/sql"

	"github.com/Iemaduddin/goweb/backend-go/internal/model"
)

type AssetRepository interface {
	CreateAsset(ctx context.Context, asset *model.Asset) error
	GetAssetByID(ctx context.Context, id int64) (*model.Asset, error)
	FindAllAssets(ctx context.Context) ([]model.Asset, error)
	UpdateAsset(ctx context.Context, asset *model.Asset) error
	DeleteAsset(ctx context.Context, id int64) error
	ActivateAsset(ctx context.Context, id int64, isActive bool) error
}

type assetRepository struct {
	db *sql.DB
}

func NewAssetRepository(db *sql.DB) AssetRepository {
	return &assetRepository{db: db}
}

func (r *assetRepository) CreateAsset(ctx context.Context, asset *model.Asset) error {
	query := `
		INSERT INTO assets (asset_code, name, category, description, is_active, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	res, err := r.db.ExecContext(
		ctx,
		query,
		asset.AssetCode,
		asset.Name,
		asset.Category,
		asset.Description,
		asset.IsActive,
		asset.CreatedAt,
	)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	asset.ID = id

	return nil
}

func (r *assetRepository) GetAssetByID(ctx context.Context, id int64) (*model.Asset, error) {
	query := `
		SELECT id, asset_code, name, category, description, is_active, created_at
		FROM assets
		WHERE id = ?
	`

	row := r.db.QueryRowContext(ctx, query, id)

	asset := &model.Asset{}

	err := row.Scan(
		&asset.ID,
		&asset.AssetCode,
		&asset.Name,
		&asset.Category,
		&asset.Description,
		&asset.IsActive,
		&asset.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (r *assetRepository) FindAllAssets(ctx context.Context) ([]model.Asset, error) {
	query := `
		SELECT id, asset_code, name, category, description, is_active, created_at
		FROM assets
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var assets []model.Asset

	for rows.Next() {
		asset := &model.Asset{}

		err := rows.Scan(
			&asset.ID,
			&asset.AssetCode,
			&asset.Name,
			&asset.Category,
			&asset.Description,
			&asset.IsActive,
			&asset.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		assets = append(assets, *asset)
	}

	return assets, nil
}

func (r *assetRepository) UpdateAsset(ctx context.Context, asset *model.Asset) error {
	query := `
		UPDATE assets
		SET asset_code = ?, name = ?, category = ?, description = ?, is_active = ?, created_at = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		asset.AssetCode,
		asset.Name,
		asset.Category,
		asset.Description,
		asset.IsActive,
		asset.CreatedAt,
		asset.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *assetRepository) DeleteAsset(ctx context.Context, id int64) error {
	query := `
		DELETE FROM assets
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *assetRepository) ActivateAsset(ctx context.Context, id int64, isActive bool) error {
	query := `
		UPDATE assets
		SET is_active = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, isActive, id)
	if err != nil {
		return err
	}

	return nil
}
