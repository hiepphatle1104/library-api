package dto

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
	Data    any    `json:"data,omitempty"`
	Token   string `json:"token,omitempty"`
}

func NewResponse(message string, ok bool) *Response {
	return &Response{
		Message: message,
		Ok:      ok,
	}
}

func (res *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&res)
}
