package server

import (
	"context"
	"fmt"
	"food-delivery-app-notification-service/internal/app/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func RDBMS(ctx context.Context) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "mysql",
		*config.Environment.DB.DBUsername+":"+
			*config.Environment.DB.DBPassword+"@("+
			*config.Environment.DB.DBHost+":"+
			*config.Environment.DB.DBPort+")/"+
			*config.Environment.DB.DBName)
	if err != nil {
		log.Panicf("%s:%s", "MySQL database connection failed...", err)
	}
	//connection pooling
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(20 * time.Second)
	log.Print("mysql database connection established successfull...")

	return db, nil
}

func LaunchServer(timeout time.Duration, routeHandler *mux.Router) *http.Server {
	fmt.Println("🚀 Launching REST Server...")

	srv := &http.Server{
		Handler:      routeHandler,
		Addr:         *config.Environment.APP.Host + ":" + *config.Environment.APP.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	return srv
}

func GracefulShutdown(ctx context.Context, srv *http.Server, timeout time.Duration, cancel func()) {
	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")

	os.Exit(0)
}
