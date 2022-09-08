package hello

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestSayHello(t *testing.T) {
	r := chi.NewRouter()

	helloSvc := &HelloServices{
		Router: r,
	}
	helloSvc.RegisterSvc()

	svr := httptest.NewServer(r)
	defer svr.Close()

	res, err := http.Get(svr.URL)
	if err != nil {
		require.Equal(t, nil, err, "Should not error")
	}

	defer res.Body.Close()
	out, err := io.ReadAll(res.Body)
	if err != nil {
		require.Equal(t, nil, err, "Should not error")
	}

	require.Equal(t, http.StatusOK, res.StatusCode, "Should return 200")
	require.Equal(t, "Hi there...", string(out), "Should return Hi there...")
}
