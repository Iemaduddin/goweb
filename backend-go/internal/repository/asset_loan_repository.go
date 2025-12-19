package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Iemaduddin/goweb/backend-go/internal/model"
)

type AssetLoanRepository interface {
	CreateAssetLoan(ctx context.Context, assetLoan *model.AssetLoan) error
	HasDateConflict(ctx context.Context, assetID int64, start, end time.Time) (bool, error)
	FindByUser(ctx context.Context, userID int64) ([]model.AssetLoan, error)
	ApproveAssetLoan(ctx context.Context, loanID, adminID int64) error
	RejectAssetLoan(ctx context.Context, loanID, adminID int64, notes string) error
	FindAllAssetLoan(ctx context.Context) ([]model.AssetLoan, error)
}

type assetLoanRepository struct {
	db *sql.DB
}

func NewAssetLoanRepository(db *sql.DB) AssetLoanRepository {
	return &assetLoanRepository{db: db}
}

func (r *assetLoanRepository) CreateAssetLoan(ctx context.Context, assetLoan *model.AssetLoan) error {
	query := `
		INSERT INTO asset_loans (user_id, asset_id, start_date, end_date, status, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	res, err := r.db.ExecContext(
		ctx,
		query,
		assetLoan.AssetID,
		assetLoan.UserID,
		assetLoan.StartDate,
		assetLoan.EndDate,
		assetLoan.Status,
		assetLoan.CreatedAt,
	)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	assetLoan.ID = id

	return nil
}

func (r *assetLoanRepository) HasDateConflict(ctx context.Context, assetID int64, start, end time.Time) (bool, error) {
	query := `
		SELECT 1
		FROM asset_loans
		WHERE asset_id = ? 
			AND  status = 'approved'
			AND NOT (
				end_date < ?
				OR start_date > ?
			)
		LIMIT 1
	`

	var exists int
	err := r.db.QueryRowContext(ctx, query, assetID, start, end).Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *assetLoanRepository) FindByUser(ctx context.Context, userID int64) ([]model.AssetLoan, error) {
	query := `
		SELECT *
		FROM asset_loans
		WHERE user_id = ?
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var assetLoans []model.AssetLoan
	for rows.Next() {
		var l model.AssetLoan
		err := rows.Scan(
			&l.ID,
			&l.UserID,
			&l.AssetID,
			&l.StartDate,
			&l.EndDate,
			&l.Status,
			&l.Notes,
			&l.ApprovedBy,
			&l.ApprovedAt,
			&l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		assetLoans = append(assetLoans, l)
	}

	return assetLoans, nil
}

func (r *assetLoanRepository) ApproveAssetLoan(ctx context.Context, loanID, adminID int64) error {
	query := `
		UPDATE asset_loans
		SET status = 'approved',
		    approved_by = ?,
		    approved_at = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, adminID, time.Now(), loanID)
	return err
}
func (r *assetLoanRepository) RejectAssetLoan(ctx context.Context, loanID, adminID int64, notes string) error {
	query := `
		UPDATE asset_loans
		SET status = 'rejected',
		    approved_by = ?,
		    approved_at = ?,
		    notes = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(ctx, query, adminID, time.Now(), notes, loanID)
	return err
}

func (r *assetLoanRepository) FindAllAssetLoan(ctx context.Context) ([]model.AssetLoan, error) {
	query := `
		SELECT *
		FROM asset_loans
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var assetLoans []model.AssetLoan
	for rows.Next() {
		var l model.AssetLoan
		err := rows.Scan(
			&l.ID,
			&l.UserID,
			&l.AssetID,
			&l.StartDate,
			&l.EndDate,
			&l.Status,
			&l.Notes,
			&l.ApprovedBy,
			&l.ApprovedAt,
			&l.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		assetLoans = append(assetLoans, l)
	}

	return assetLoans, nil
}
