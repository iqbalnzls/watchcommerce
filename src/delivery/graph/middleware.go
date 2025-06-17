package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	appContext "github.com/iqbalnzls/watchcommerce/src/shared/app_context"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
	"github.com/iqbalnzls/watchcommerce/src/shared/logger"
)

func SetupMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appCtx := appContext.NewAppContext(&logger.Log{
			XID:         uuid.New().String(),
			Time:        time.Now(),
			Path:        r.URL.Path,
			ServiceName: constant.AppName,
			Version:     constant.AppVersion,
			Header:      r.Header,
			IP:          r.RemoteAddr,
		})

		appCtx.Logger.IncomingRequest()

		cont := context.WithValue(r.Context(), constant.AppContext, appCtx)

		next.ServeHTTP(w, r.WithContext(cont))
	})
}
