package authentication

import (
	"context"
	"fmt"
	"net/http"

	"go_grpc_boileplate/common/constant"
	"go_grpc_boileplate/common/http_response"
	"go_grpc_boileplate/common/middlewares/authorization"
	"go_grpc_boileplate/configs"

	"github.com/bytedance/sonic"
	"github.com/golang-jwt/jwt/v4"
)

func Authentication() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authToken := r.Header.Get("Authorization")

			responseError := http_response.HttpResponse{
				Status:  http.StatusForbidden,
				Message: constant.MSG_FORBIDDEN_ACCESS,
				Data:    nil,
			}

			if authToken == "" {
				http_response.New(w, responseError).Send()
				return
			}

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
				}

				return []byte(configs.Config.JWT.SecretKey), nil
			})
			if err != nil {
				http_response.New(w, responseError).Send()
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				http_response.New(w, responseError).Send()
				return
			}
			b, err := sonic.Marshal(claims)
			if err != nil {
				http_response.New(w, responseError).Send()
				return
			}

			var userInfo authorization.UserInfo

			err = sonic.Unmarshal(b, &userInfo)
			if err != nil {
				http_response.New(w, responseError).Send()
				return
			}

			ctx := context.WithValue(r.Context(), "userInfo", userInfo)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
