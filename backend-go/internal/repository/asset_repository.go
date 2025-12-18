package repository

type AssetRepository interface {
	GetAllAssets() ([]Asset, error)
	GetAssetByID(id int64) (Asset, error)
	CreateAsset(asset Asset) error
	UpdateAsset(asset Asset) error
	DeleteAsset(id int64) error
}
