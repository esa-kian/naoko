package main

import (
	"encoding/json"
	"log"
	"naoko/services"
	"net/http"

	"github.com/gorilla/mux"
)

type ConnectRequest struct {
	DBType   string `json:"db_type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

// Handler functions - these will be fleshed out later
func ConnectDatabase(w http.ResponseWriter, r *http.Request) {
	var req ConnectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Attempt to connect using the provided details
	if err := services.ConnectDatabase(req.DBType, req.Host, req.Port, req.User, req.Password, req.DBName); err != nil {
		http.Error(w, "Failed to connect to database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database connected successfully"))
}

func GenerateReport(w http.ResponseWriter, r *http.Request) {
	// Placeholder for report generation logic
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Report generated successfully"))
}

func main() {
	// Initialize the Gorilla Mux router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/connect", ConnectDatabase).Methods("POST")
	r.HandleFunc("/report", GenerateReport).Methods("POST")

	// Set up the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
