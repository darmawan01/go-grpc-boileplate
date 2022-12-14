package http_response

import (
	"net/http"

	"go_grpc_boileplate/common/constant"

	"github.com/bytedance/sonic"
)

type HttpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Meta    *Meta  `json:"meta,omitempty"`
}

type Meta struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}

type response struct {
	response HttpResponse
	writer   http.ResponseWriter
}

func New(writer http.ResponseWriter, res HttpResponse) *response {
	return &response{
		response: res,
		writer:   writer,
	}
}

func (h *response) Send() {
	h.writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	res, err := sonic.Marshal(h.response)
	if err != nil {
		h.writer.WriteHeader(http.StatusForbidden)
		h.writer.Write([]byte(constant.MSG_FORBIDDEN_ACCESS))
		return
	}

	h.writer.WriteHeader(h.response.Status)
	h.writer.Write(res)
}
