package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Message:   "Halo dari Backend Go (Golang)!",
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    "success",
	}

	json.NewEncoder(w).Encode(res)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
