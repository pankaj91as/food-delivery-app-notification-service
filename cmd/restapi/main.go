package main

import (
	"context"
	"flag"
	"fmt"
	"food-delivery-app-notification-service/internal/app/controller"
	"food-delivery-app-notification-service/internal/app/repository"
	"food-delivery-app-notification-service/internal/app/router"
	"food-delivery-app-notification-service/internal/app/service"
	"food-delivery-app-notification-service/server"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "graceful-timeout", 1*time.Minute, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// Panic Recover Functionality
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Recovered in restapi: ", r)
		}
	}()

	dbConnection, _ := server.RDBMS(ctx)
	restRepo := repository.NewRepoInit(dbConnection)
	restService := service.NewRestService(restRepo)
	restController := controller.NewRestController(restService)

	r := router.NewRouter(mux.NewRouter(), restController)
	restSrv := server.LaunchServer(timeout, r.RestHandler())

	//during shutdown this function will execute to release the resources
	allCancel := func() {
		cancel()
	}

	// Graceful Shutdown Server
	server.GracefulShutdown(ctx, restSrv, timeout, allCancel)
}
