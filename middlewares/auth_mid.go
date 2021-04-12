package middlewares

import (
	"net/http"

	"../jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		_, err := jwt.ProcessToken(header)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		next.ServeHTTP(w, r)
	}
}