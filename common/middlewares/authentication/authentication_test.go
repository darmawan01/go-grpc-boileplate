package authentication

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go_grpc_boileplate/common/constant"
	"go_grpc_boileplate/common/http_response"
	"go_grpc_boileplate/common/test"
	"go_grpc_boileplate/configs"

	"github.com/bytedance/sonic"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

func TestAuthentication(t *testing.T) {
	r := chi.NewRouter()
	key := "secretKey"
	configs.Config.JWT.SecretKey = key
	responseError := http_response.HttpResponse{
		Status:  http.StatusForbidden,
		Message: constant.MSG_FORBIDDEN_ACCESS,
		Data:    nil,
	}

	b, err := sonic.Marshal(responseError)
	if err != nil {
		t.Fatal(err)
	}

	r.Use(Authentication())
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	// sending unauthorized requests
	if status, resp := test.TestRequest(t, ts, "GET", "/", nil, nil); status != http.StatusForbidden || resp != string(b) {
		t.Fatalf(resp)
	}

	// sending wrong key
	h := http.Header{}
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	token, err := jwtToken.SignedString([]byte("wrong"))
	if err != nil {
		t.Fatal(err)
	}
	h.Set("Authorization", token)
	if status, resp := test.TestRequest(t, ts, "GET", "/", h, nil); status != http.StatusForbidden || resp != string(b) {
		t.Fatalf(resp)
	}

	// sending wrong jwt token
	h.Set("Authorization", "asdf")
	if status, resp := test.TestRequest(t, ts, "GET", "/", h, nil); status != http.StatusForbidden || resp != string(b) {
		t.Fatalf(resp)
	}

	// sending authorized requests
	jwtToken = jwt.New(jwt.SigningMethodHS256)
	token, err = jwtToken.SignedString([]byte(key))
	if err != nil {
		t.Fatal(err)
	}
	h.Set("Authorization", token)
	if status, resp := test.TestRequest(t, ts, "GET", "/", h, nil); status != 200 || resp != "welcome" {
		t.Fatalf(resp)
	}
}
