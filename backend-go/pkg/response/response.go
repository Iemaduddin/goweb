package response

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteJson(w http.ResponseWriter, statusCode int, payload JSONResponse) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(payload)

}
