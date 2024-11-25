package controller

import (
	"encoding/json"
	"food-delivery-app-notification-service/pkg/model"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	data := model.SampleData{
		Data: "test",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
