package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Iemaduddin/goweb/backend-go/internal/model"
	"github.com/Iemaduddin/goweb/backend-go/internal/service"
)

type AssetLoanHandler struct {
	service service.AssetLoanService
}

func NewAssetLoanHandler(s service.AssetLoanService) *AssetLoanHandler {
	return &AssetLoanHandler{service: s}
}

type requestLoanDTO struct {
	AssetID   int64     `json:"asset_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (h *AssetLoanHandler) RequestLoan(w http.ResponseWriter, r *http.Request) {
	var req requestLoanDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	loan := &model.AssetLoan{
		UserID:    getUserIDFromContext(r),
		AssetID:   req.AssetID,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	if err := h.service.RequestLoan(r.Context(), loan); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loan)
}
