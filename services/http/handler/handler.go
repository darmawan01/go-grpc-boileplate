package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HandlerServices struct {
	Router *chi.Mux
}

func (svc *HandlerServices) RegisterSvc() {
	svc.Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("404 Route not found"))
	})
	svc.Router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("405 Mehotd not allowed"))
	})
}
