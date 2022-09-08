package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

	res, err := http.Get(svr.URL + "/opapa")
	if err != nil {
		require.Equal(t, nil, err, "Should not error")
	}
	defer res.Body.Close()

	require.Equal(t, http.StatusNotFound, res.StatusCode, "Should return 404")

	// Test not allowed method
	res, err = http.Post(svr.URL, "application/json", nil)
	if err != nil {
		require.Equal(t, nil, err, "Should not error")
	}
	defer res.Body.Close()

	require.Equal(t, http.StatusMethodNotAllowed, res.StatusCode, "Should return 405")
}
