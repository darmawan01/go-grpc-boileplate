package http_responses

import (
	"encoding/json"
	"go_grpc_boileplate/common/constant"
	"net/http"
)

type HttpResponses struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Meta    any    `json:"meta,omitempty"`
}

type responses struct {
	response HttpResponses
	writer   http.ResponseWriter
}

func New(writer http.ResponseWriter, response HttpResponses) *responses {
	return &responses{
		response: response,
		writer:   writer,
	}
}

func (h *responses) Send() {
	h.writer.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(h.response)
	if err != nil {
		h.writer.WriteHeader(http.StatusForbidden)
		h.writer.Write([]byte(constant.MSG_FORBIDDEN_ACCESS))
		return
	}

	h.writer.WriteHeader(h.response.Status)
	h.writer.Write(res)
}
