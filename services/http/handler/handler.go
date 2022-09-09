package handler

import (
	"net/http"

	"go_grpc_boileplate/common/constant"
	"go_grpc_boileplate/common/http_response"

	"github.com/go-chi/chi/v5"
)

type HandlerServices struct {
	Router *chi.Mux
}

func (svc *HandlerServices) RegisterSvc() {
	svc.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http_response.New(w, http_response.HttpResponse{
			Status:  http.StatusNotFound,
			Message: constant.MSG_NOT_FOUND,
		}).Send()
	})

	svc.Router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		http_response.New(w, http_response.HttpResponse{
			Status:  http.StatusMethodNotAllowed,
			Message: constant.MSG_NOT_FOUND,
		}).Send()
	})
}
