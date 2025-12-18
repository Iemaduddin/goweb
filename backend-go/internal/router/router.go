package router

import (
	"net/http"

	"github.com/Iemaduddin/goweb/backend-go/internal/handler"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", http.HandlerFunc(handler.Home))

	return mux
}
