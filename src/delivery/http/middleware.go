package http

import (
	"context"
	"net/http"

	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	"github.com/iqbalnzls/watchcommerce/src/pkg/logger"
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

			logger := logger.NewLogger(&logger.Log{
				Path:        r.URL.Path,
				ServiceName: constant.AppName,
				Version:     constant.AppVersion,
				Header:      r.Header,
				IP:          r.RemoteAddr,
			})

			logger.IncomingRequest()

			cont := context.WithValue(r.Context(), constant.AppContext, logger)

			next.ServeHTTP(w, r.WithContext(cont))
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
