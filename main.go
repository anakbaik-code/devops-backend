package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
}

func main() {
	// Endpoint API: /api/hello
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		res := Response{
			Message:   "Halo dari Backend Go (Golang)!",
			Timestamp: time.Now().Format(time.RFC3339),
			Status:    "success",
		}

		json.NewEncoder(w).Encode(res)
	})

	// Endpoint health check
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println("Server Go berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
