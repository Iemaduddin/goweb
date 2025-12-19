package model

import "time"

type AssetLoan struct {
	ID      int64 `json:"id"`
	UserID  int64 `json:"user_id"`
	AssetID int64 `json:"asset_id"`

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`

	Status string `json:"status"`
	Notes  string `json:"notes"`

	ApprovedBy *int64     `json:"approved_by,omitempty"`
	ApprovedAt *time.Time `json:"approved_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}
