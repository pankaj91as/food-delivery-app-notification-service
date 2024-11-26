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
	r.router.HandleFunc("/", r.restController.Handler).Methods(http.MethodGet)
	return r.router
}
