package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Service   string `json:"service"`
	Status    string `json:"status"`
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/metrics", metricsHandler)

	log.Println("InfraWatch API running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("InfraWatch API is running"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Service:   "infrawatch-api",
		Status:    "healthy",
		Version:   "1.0.0",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"version":"1.0.0","environment":"development"}`))
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"service":        "infrawatch-api",
		"uptime_status":  "running",
		"request_status": "ok",
		"cpu_usage":      "simulated",
		"memory_usage":   "simulated",
		"timestamp":      time.Now().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(response)
}