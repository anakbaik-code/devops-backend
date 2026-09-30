package main

import (
	"fmt"
	"go-backend/handler"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// Endpoint API: /api/hello
	http.HandleFunc("/api/hello", handler.HelloHandler)

	// Endpoint health check
	http.HandleFunc("/api/health", handler.HealthHandler)

	port := os.Getenv("BACKEND_PORT")
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server Berjalan Di Port %s ... ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
