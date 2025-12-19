package service

import (
	"context"
	"errors"

	"github.com/Iemaduddin/goweb/backend-go/internal/model"
	"github.com/Iemaduddin/goweb/backend-go/internal/repository"
)

type AssetLoanService interface {
	RequestLoan(ctx context.Context, loan *model.AssetLoan) error
	GetLoansByUser(ctx context.Context, userID int64) ([]model.AssetLoan, error)
	GetAllLoans(ctx context.Context) ([]model.AssetLoan, error)
	ApproveLoan(ctx context.Context, loanID, adminID int64) error
	RejectLoan(ctx context.Context, loanID, adminID int64, notes string) error
}

type assetLoanService struct {
	assetLoanRepo repository.AssetLoanRepository
	assetRepo     repository.AssetRepository
}

func NewAssetLoanService(repo repository.AssetLoanRepository, assetRepo repository.AssetRepository) AssetLoanService {
	return &assetLoanService{
		assetLoanRepo: repo,
		assetRepo:     assetRepo,
	}
}

func (s *assetLoanService) RequestLoan(ctx context.Context, loan *model.AssetLoan) error {
	// validasi tanggal
	if loan.StartDate.After(loan.EndDate) {
		return errors.New("start date must be before end date")
	}

	// validasi asset
	asset, err := s.assetRepo.GetAssetByID(ctx, loan.AssetID)
	if err != nil {
		return err
	}

	if asset == nil || !asset.IsActive {
		return errors.New("asset not available")
	}

	// validasi konflik tanggal
	conflict, err := s.assetLoanRepo.HasDateConflict(ctx, loan.AssetID, loan.StartDate, loan.EndDate)
	if err != nil {
		return err
	}

	if conflict {
		return errors.New("asset is already loaned on selected date")
	}

	// simpan peminjaman
	return s.assetLoanRepo.CreateAssetLoan(ctx, loan)
}

func (s *assetLoanService) GetLoansByUser(ctx context.Context, userID int64) ([]model.AssetLoan, error) {
	return s.assetLoanRepo.FindByUser(ctx, userID)
}

func (s *assetLoanService) GetAllLoans(ctx context.Context) ([]model.AssetLoan, error) {
	return s.assetLoanRepo.FindAllAssetLoan(ctx)
}

func (s *assetLoanService) ApproveLoan(ctx context.Context, loanID, adminID int64) error {
	return s.assetLoanRepo.ApproveAssetLoan(ctx, loanID, adminID)
}

func (s *assetLoanService) RejectLoan(ctx context.Context, loanID, adminID int64, notes string) error {
	return s.assetLoanRepo.RejectAssetLoan(ctx, loanID, adminID, notes)
}
