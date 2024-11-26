package router

import (
	"food-delivery-app-notification-service/internal/app/controller"
	"food-delivery-app-notification-service/internal/app/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type IRouter interface {
	RestHandler() *mux.Router
}

type Router struct {
	router         *mux.Router
	restController controller.IRestController
}

func NewRouter(router *mux.Router, restController controller.IRestController) IRouter {
	router.Use(middleware.LoggingMiddleware)

	return &Router{
		router:         router,
		restController: restController,
	}
}

func (r *Router) RestHandler() *mux.Router {
	r.router.HandleFunc("/api/v1/orders", r.restController.GetOrders).Methods(http.MethodGet)
	r.router.HandleFunc("/api/v1/orders/{orderid}", r.restController.GetOrdersByID).Methods(http.MethodGet)
	r.router.HandleFunc("/api/v1/orders/{orderid}", r.restController.UpdateOrderByID).Methods(http.MethodPatch)
	r.router.HandleFunc("/api/v1/customers", r.restController.GetCustomers).Methods(http.MethodGet)
	return r.router
}
