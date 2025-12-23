package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Iemaduddin/goweb/backend-go/internal/service"
)

type AssetHandler struct {
	service service.AssetService
}

func NewAssetHandler(s service.AssetService) *AssetHandler {
	return &AssetHandler{service: s}
}

type requestAssetDTO struct {
	ID        int64  `json:"id"`
	AssetCode string `json:"asset_code"`
	AssetType string `json:"asset_type"`
	Location  string `json:"location"`
	Quantity  int    `json:"quantity"`
	IsActive  bool   `json:"is_active"`
}

func (h *AssetHandler) CreateAsset(w http.ResponseWriter, r *http.Request) {
	var req requestAssetDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	asset := &model.Asset{
		ID:        req.ID,
		AssetCode: req.AssetCode,
		AssetType: req.AssetType,
		Location:  req.Location,
		Quantity:  req.Quantity,
		IsActive:  req.IsActive,
	}

	if err := h.service.CreateAsset(r.Context(), asset); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(asset)
}

func (h *AssetHandler) GetAssetByID(w http.ResponseWriter, r *http.Request) {
	var req requestAssetDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetAssetByID(r.Context(), req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
