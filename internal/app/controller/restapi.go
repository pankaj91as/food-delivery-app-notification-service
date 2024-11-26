package controller

import (
	"food-delivery-app-notification-service/internal/app/service"
	"net/http"
)

type IRestController interface {
	GetOrders(w http.ResponseWriter, r *http.Request)
	GetOrdersByID(w http.ResponseWriter, r *http.Request)
	UpdateOrderByID(w http.ResponseWriter, r *http.Request)

	GetCustomers(w http.ResponseWriter, r *http.Request)
}

type RestController struct {
	restService service.IRestService
}

func NewRestController(restService service.IRestService) IRestController {
	return &RestController{
		restService: restService,
	}
}
