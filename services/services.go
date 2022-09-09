package services

import (
	"go_grpc_boileplate/services/http/handler"
	"go_grpc_boileplate/services/http/hello"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Services struct {
	Router *chi.Mux
	DB     *gorm.DB
}

func (svc *Services) Resgiters() {
	helloSvc := hello.HelloServices{
		Router: svc.Router,
		DB:     svc.DB,
	}
	helloSvc.RegisterSvc()

	handlerSvc := handler.HandlerServices{
		Router: svc.Router,
	}
	handlerSvc.RegisterSvc()
}
