package http_response

import (
	"encoding/json"
	"go_grpc_boileplate/common/constant"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func newServer() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		New(w,
			HttpResponse{
				Status:  http.StatusOK,
				Message: "success",
			},
		).Send()
	}))

	router.HandleFunc("/forbidden", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var unsupportedValues = []any{
			math.NaN(),
		}

		New(w,
			HttpResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data:    unsupportedValues,
			},
		).Send()
	}))

	router.HandleFunc("/with-meta", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		New(w,
			HttpResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data:    []string{"Hello", "World"},
				Meta: &Meta{
					Page:      1,
					TotalPage: 1,
					TotalData: 2,
				},
			},
		).Send()
	}))

	return router
}

func TestHttpResponse(t *testing.T) {
	svr := httptest.NewServer(newServer())
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

	var resp HttpResponse
	if err := json.Unmarshal(out, &resp); err != nil {
		require.Equal(t, nil, err, "Should not error")
	}

	require.Equal(t, http.StatusOK, resp.Status, "Should return 200")
	require.Equal(t, "success", resp.Message, "Should return success")
	require.Equal(t, (*Meta)(nil), resp.Meta, "Should return nil")

	// Test on failed to marshal data
	res, err = http.Get(svr.URL + "/forbidden")
	if err != nil {
		require.Equal(t, nil, err, "Should not error")
	}

	defer res.Body.Close()
	out, err = io.ReadAll(res.Body)
	if err != nil {
		require.Equal(t, nil, err, "Should not error")
	}

	require.Equal(t, http.StatusForbidden, res.StatusCode, "Should return 403")
	require.Equal(t, constant.MSG_FORBIDDEN_ACCESS, string(out), "Should return forbidden access")

	// Test get meta
	res, err = http.Get(svr.URL + "/with-meta")
	if err != nil {
		require.Equal(t, nil, err, "Should not error")
	}

	defer res.Body.Close()
	out, err = io.ReadAll(res.Body)
	if err != nil {
		require.Equal(t, nil, err, "Should not error")
	}

	if err := json.Unmarshal(out, &resp); err != nil {
		require.Equal(t, nil, err, "Should not error")
	}

	require.Equal(t, http.StatusOK, resp.Status, "Should return 200")
	require.Equal(t, "success", resp.Message, "Should return success")
	require.NotEqual(t, nil, resp.Meta, "Should not return nil")
	require.Equal(t, 1, resp.Meta.Page, "Should return page 1")
	require.Equal(t, 1, resp.Meta.TotalPage, "Should return total page 1")
	require.Equal(t, 2, resp.Meta.TotalData, "Should return total data 2")
}
