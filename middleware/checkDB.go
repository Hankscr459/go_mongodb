package middleware

import (
	"net/http"
	"twitter/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Connection lost Data", 500)
		}
		next.ServeHTTP(w, r)
	}
}
