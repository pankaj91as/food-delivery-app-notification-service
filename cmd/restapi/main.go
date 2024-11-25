package main

import (
	"food-delivery-app-notification-service/internal/app/controller"
	"food-delivery-app-notification-service/internal/app/middleware"
	"food-delivery-app-notification-service/internal/app/server"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Handler)
	r.Use(middleware.LoggingMiddleware)
	server.LaunchServer(r)
}
