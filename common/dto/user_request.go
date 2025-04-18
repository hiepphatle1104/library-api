package dto

import (
	"encoding/json"
	"net/http"
)

func GetRequestData[T any](r *http.Request) *T {
	var data T
	json.NewDecoder(r.Body).Decode(&data)
	defer r.Body.Close()

	return &data
}
