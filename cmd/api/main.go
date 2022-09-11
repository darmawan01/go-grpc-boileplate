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
	dbConn     *gorm.DB
	grpcServer *grpc.Server
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	configs.LoadFromEnv()

	var err error
	if dbConn, err = db.Open(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	/*
		serverCtx, serverStopCtx := context.WithCancel(context.Background())
		defer serverStopCtx()
		// load config from vault and auto reload
		vaultAddress := os.Getenv("VAULT_ADDRESS")
		vaultToken := os.Getenv("VAULT_TOKEN")
		configs.LoadFromVault(vaultAddress, vaultToken)
		var err error
		if dbConn, err = db.Open(); err != nil {
			log.Fatal(err)
		}

		// TODO: maybe we can use endpoint to trigger this and load all services like dbConn or redisConn, but dbConn or redisConn must public var on their pkg
		go func() {
			t := time.NewTicker(5 * time.Minute)
			defer t.Stop()
			for {
				select {
				case <-serverCtx.Done():
					return
				case <-t.C:
					configs.LoadFromVault(vaultAddress, vaultToken)
				}
			}
		}()
	*/

	log.Println("Start server...")

	// Serve HTTP
	httpServer := serveHttp()

	// Serve GRPC
	go serveGRPC()

	// Server run context
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		// Run the server
		log.Printf("Server is running on port %s \n", configs.Config.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for signal close
	<-sig

	// serverStopCtx()
	fmt.Println()
	log.Println("closing grpc server")
	grpcServer.GracefulStop()

	// Shutdown signal with grace period of 30 seconds
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Trigger graceful shutdown
	log.Println("closing http server")
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}

func serveHttp() *http.Server {
	r := chi.NewRouter()

	// Middleware
	if !configs.Config.IsProduction() {
		r.Use(middleware.Logger)
	}

	// Register services
	service := services.Services{Router: r, DB: dbConn}
	service.Register()

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", configs.Config.Port),
		Handler: r,
	}
}

func serveGRPC() {
	grpcServer = grpc.NewServer()

	pb.RegisterHelloServicesServer(grpcServer, &pb.HelloGrpcServices{})

	listener, err := net.Listen("tcp", ":9002")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(listener)
}
