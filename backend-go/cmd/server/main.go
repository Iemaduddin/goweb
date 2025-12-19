package main

import (
	"log"
	"net/http"

	"github.com/Iemaduddin/goweb/backend-go/internal/config"
	"github.com/Iemaduddin/goweb/backend-go/internal/database"
	"github.com/Iemaduddin/goweb/backend-go/internal/handler"
	"github.com/Iemaduddin/goweb/backend-go/internal/repository"
	"github.com/Iemaduddin/goweb/backend-go/internal/router"
	"github.com/Iemaduddin/goweb/backend-go/internal/service"
)

func main() {
	cfg := config.Load()
	db, err := database.MySQLDB(cfg)
	if err != nil {
		log.Fatal("Terjadi kesalahan: ", err)
	}
	defer db.Close()

	assetRepo := repository.NewAssetRepository(db)
	loanRepo := repository.NewAssetLoanRepository(db)

	loanService := service.NewAssetLoanService(loanRepo, assetRepo)
	loanHandler := handler.NewAssetLoanHandler(loanService)

	router := router.NewRouter(loanHandler)

	log.Println("server running on :8080")
	http.ListenAndServe(":8080", router)
}
