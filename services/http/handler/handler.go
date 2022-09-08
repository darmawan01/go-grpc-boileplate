package handler

import (
	"net/http"

	"go_grpc_boileplate/common/constant"
	"go_grpc_boileplate/common/http_responses"

	"github.com/go-chi/chi/v5"
)

type HandlerServices struct {
	Router *chi.Mux
}

func (svc *HandlerServices) RegisterSvc() {
	svc.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http_responses.New(w, http_responses.HttpResponses{
			Status:  http.StatusNotFound,
			Message: constant.MSG_NOT_FOUND,
		}).Send()
	})

	svc.Router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		http_responses.New(w, http_responses.HttpResponses{
			Status:  http.StatusMethodNotAllowed,
			Message: constant.MSG_NOT_FOUND,
		}).Send()
	})
}
