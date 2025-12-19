package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Iemaduddin/goweb/backend-go/internal/model"
	"github.com/Iemaduddin/goweb/backend-go/internal/service"
	"github.com/go-chi/chi/v5"
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(loan)
}

func (h *AssetLoanHandler) GetLoansByUser(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	loans, err := h.service.GetLoansByUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(loans)
}

func (h *AssetLoanHandler) GetAllLoans(w http.ResponseWriter, r *http.Request) {
	loans, err := h.service.GetAllLoans(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(loans)
}

func (h *AssetLoanHandler) ApproveLoan(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	loanID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid loan ID", http.StatusBadRequest)
		return
	}

	if err := h.service.ApproveLoan(r.Context(), loanID, getUserIDFromContext(r)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *AssetLoanHandler) RejectLoan(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	loanID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid loan ID", http.StatusBadRequest)
		return
	}

	if err := h.service.RejectLoan(r.Context(), loanID, getUserIDFromContext(r)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
