package main

import (
	"flag"
	"food-delivery-app-notification-service/internal/app/controller"
	"food-delivery-app-notification-service/internal/app/middleware"
	"food-delivery-app-notification-service/internal/app/server"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", controller.Handler)
	r.Use(middleware.LoggingMiddleware)
	server.LaunchServer(timeout, r)
}
