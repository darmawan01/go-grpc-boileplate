package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go_grpc_boileplate/common/db"
	"go_grpc_boileplate/configs"
	"go_grpc_boileplate/services"

	pb "go_grpc_boileplate/services/grpc/hello"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	conf   *configs.Configs
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf = configs.LoadFromEnv()

	conn := &db.DBConn{
		Info:       conf.DB,
		SilentMode: conf.IsProduction(),
	}

	var err error
	if dbConn, err = conn.Open(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Start server...")

	// Serve HTTP
	httpServer := serveHttp()

	// Serve GRPC
	go serveGRPC()

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 5*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			log.Fatal(err)
		}

		serverStopCtx()
		// grpcServer.Stop()
	}()

	// Run the server
	log.Printf("Server is running on port %s \n", conf.Port)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

func serveHttp() *http.Server {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Register services
	service := services.Services{Router: r, DB: dbConn}
	service.Register()

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", conf.Port),
		Handler: r,
	}
}

func serveGRPC() *grpc.Server {
	grpcServer := grpc.NewServer()

	pb.RegisterHelloServicesServer(grpcServer, &pb.HelloGrpcServices{})

	listener, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listener)

	return grpcServer
}
