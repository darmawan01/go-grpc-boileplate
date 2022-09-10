package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go_grpc_boileplate/common/test"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/stretchr/testify/require"
)

func newServer() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedMethods: []string{"GET"},
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	helloSvc := &HandlerServices{
		Router: r,
	}
	helloSvc.RegisterSvc()

	return r
}

func TestHandler(t *testing.T) {
	svr := httptest.NewServer(newServer())
	defer svr.Close()

	statusCode, _ := test.TestRequest(t, svr, "GET", "/opapa", nil, nil)

	require.Equal(t, http.StatusNotFound, statusCode, "Should return 404")

	// Test not allowed method
	h := http.Header{}
	h.Add("Content-Type", "application/json")
	statusCode, _ = test.TestRequest(t, svr, "POST", "/", h, nil)

	require.Equal(t, http.StatusMethodNotAllowed, statusCode, "Should return 405")
}
