package http

import "net/http"

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func SetupMiddleware() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			//handler CORS
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Content-Type", "application/json")

			next.ServeHTTP(w, r)
		}
	}
}
