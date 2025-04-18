package auth

import (
	"library-api/common/dto"
	"net/http"
)

func VerifyPostData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.Header().Set("Content-Type", "application/json")
			res := dto.NewResponse("invalid data", false)
			res.Send(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
