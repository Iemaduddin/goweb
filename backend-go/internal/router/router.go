package router

import (
	"net/http"

	"github.com/Iemaduddin/goweb/backend-go/internal/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(assetLoanHandler *handler.AssetLoanHandler) http.Handler {
	r := chi.NewRouter()

	r.Post("/api/v1/asset-loans", assetLoanHandler.RequestLoan)
	r.Get("/api/v1/asset-loans/me", assetLoanHandler.GetLoansByUser)
	r.Patch("/api/v1/asset-loans/{id}/approve", assetLoanHandler.ApproveLoan)
	r.Patch("/api/v1/asset-loans/{id}/reject", assetLoanHandler.RejectLoan)

	return r
}
