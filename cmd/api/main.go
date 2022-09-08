package main

import (
	"context"
	"go_grpc_boileplate/services/http/handler"
	"go_grpc_boileplate/services/http/hello"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {

}

func main() {
	// The HTTP Server
	server := &http.Server{
		Addr:    ":3033",
		Handler: registerServices(),
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

func registerServices() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Register services
	helloSvc := hello.HelloServices{
		Router: r,
	}
	helloSvc.RegisterSvc()

	handlerSvc := handler.HandlerServices{
		Router: r,
	}
	handlerSvc.RegisterSvc()
	// End of registered services

	return r
}
