package hello

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HelloServices struct {
	Router *chi.Mux
}

func (svc *HelloServices) RegisterSvc() {
	svc.Router.Get("/", svc.sayHello)
}

func (svc *HelloServices) sayHello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hi there..."))
}
