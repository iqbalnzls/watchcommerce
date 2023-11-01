package http

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/iqbalnzls/watchcommerce/src/pkg/app_context"
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

			appCtx := app_context.NewAppContext(&logger.Log{
				XID:         uuid.New().String(),
				Time:        time.Now().String(),
				Path:        r.URL.Path,
				ServiceName: constant.AppName,
				Version:     constant.AppVersion,
				Header:      r.Header,
				IP:          r.RemoteAddr,
			})

			appCtx.Logger.IncomingRequest()

			cont := context.WithValue(r.Context(), constant.AppContext, appCtx)

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
