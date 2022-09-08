package http_responses

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
			HttpResponses{
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
			HttpResponses{
				Status:  http.StatusOK,
				Message: "success",
				Data:    unsupportedValues,
			},
		).Send()
	}))

	return router
}

func TestHttpResponses(t *testing.T) {
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

	var resp HttpResponses
	if err := json.Unmarshal(out, &resp); err != nil {
		require.Equal(t, nil, err, "Should not error")
	}

	require.Equal(t, http.StatusOK, resp.Status, "Should return 200")
	require.Equal(t, "success", resp.Message, "Should return success")

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
}
