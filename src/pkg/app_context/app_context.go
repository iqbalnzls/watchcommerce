package app_context

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/iqbalnzls/watchcommerce/src/pkg/constant"
	"github.com/iqbalnzls/watchcommerce/src/pkg/logger"
)

type AppContext struct {
	Logger logger.Logger
}

func NewAppContext(log *logger.Log) *AppContext {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	return &AppContext{
		Logger: &logger.Log{
			XID:         log.XID,
			Path:        log.Path,
			Header:      log.Header,
			Time:        log.Time,
			ServiceName: log.ServiceName,
			Version:     log.Version,
			ZapLog:      zapLogger,
		},
	}
}

func ParsingAppContext(r *http.Request) *AppContext {
	ctx := r.Context()
	appCtx, ok := ctx.Value(constant.AppContext).(*AppContext)
	if !ok {
		panic("please set app context")
	}

	return appCtx
}
