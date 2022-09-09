package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go_grpc_boileplate/common/test"

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

	status, resp := test.TestRequest(t, svr, "GET", "/", nil, nil)

	require.Equal(t, http.StatusOK, status, "Should return 200")
	require.Equal(t, "Hi there...", string(resp), "Should return Hi there...")
}
