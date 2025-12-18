package main

import (
	"log"
	"net/http"

	"github.com/Iemaduddin/goweb/backend-go/internal/router"
)

func main() {
	r := router.NewRouter()

	log.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
