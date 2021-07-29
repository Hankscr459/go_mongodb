package middleware

import (
	"net/http"
	"twitter/routers"
)

func VaildJWT(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProccessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error on Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
