package controller

import (
	"encoding/json"
	"food-delivery-app-notification-service/internal/app/service"
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
	ctx := r.Context()
	data := c.restService.GetOrders(ctx)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
