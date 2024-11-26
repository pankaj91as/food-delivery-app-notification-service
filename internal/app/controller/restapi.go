package controller

import (
	"encoding/json"
	"food-delivery-app-notification-service/internal/app/service"
	"food-delivery-app-notification-service/pkg/model"
	"net/http"
)

type IRestController interface {
	Handler(w http.ResponseWriter, r *http.Request)
}

type RestController struct {
	restService service.IRestService
}

func NewRestController(restService service.IRestService) IRestController {
	return &RestController{
		restService: restService,
	}
}

func (c *RestController) Handler(w http.ResponseWriter, r *http.Request) {
	data := model.SampleData{
		Data: "test",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
