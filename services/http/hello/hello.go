package hello

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type HelloServices struct {
	Router *chi.Mux
	DB     *gorm.DB
}

func (svc *HelloServices) RegisterSvc() {
	svc.Router.Get("/", svc.sayHello)
}

func (svc *HelloServices) sayHello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hi there..."))
}
