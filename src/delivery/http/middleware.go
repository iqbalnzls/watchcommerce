package http

import (
	"net/http"

	"github.com/iqbalnzls/watchcommerce/src/pkg/utils"
)

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func SetupMiddleware() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if !isAuthSwaggerValid(r) {
				http.Error(w, "Invalid Authorization", http.StatusForbidden)
				return
			}

			//handle CORS
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Content-Type", "application/json")

			next.ServeHTTP(w, r)
		}
	}
}

func isAuthSwaggerValid(r *http.Request) bool {
	auth := r.Header.Get("Authorization-Swagger")
	if !utils.IsEmptyString(auth) && auth != "asdjkhNasdb90834aSD" {
		return false
	}

	return true
}
