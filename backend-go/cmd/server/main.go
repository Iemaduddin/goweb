package main

import (
	"log"
	"net/http"

	"github.com/Iemaduddin/goweb/backend-go/internal/config"
	"github.com/Iemaduddin/goweb/backend-go/internal/database"
	"github.com/Iemaduddin/goweb/backend-go/internal/router"
)

func main() {
	r := router.NewRouter()
	cfg := config.Load()
	db, err := database.MySQLDB(cfg)

	if err != nil {
		log.Fatal("Terjadi kesalahan: ", err)
	}
	defer db.Close()

	log.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
