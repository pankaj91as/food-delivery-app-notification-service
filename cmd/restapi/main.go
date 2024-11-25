package main

import (
	"encoding/json"
	"food-delivery-app-notification-service/internal/app/middleware"
	"food-delivery-app-notification-service/internal/app/server"
	"net/http"

	"github.com/gorilla/mux"
)

type SampleData struct {
	Data string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler)
	r.Use(middleware.LoggingMiddleware)
	server.LaunchServer(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	data := SampleData{
		Data: "test",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
