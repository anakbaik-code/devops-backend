package main

import (
	"fmt"
	"go-backend/handler"
	"log"
	"net/http"
)

func main() {
	// Endpoint API: /api/hello
	http.HandleFunc("/api/hello", handler.HelloHandler)

	// Endpoint health check
	http.HandleFunc("/api/health",handler.HealthHandler)

	fmt.Println("Server Go berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
