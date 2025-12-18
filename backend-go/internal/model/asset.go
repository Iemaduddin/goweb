package model

import "time"

type Asset struct {
	ID          int64     `json:"id"`
	AssetCode   string    `json:"asset_code"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}
