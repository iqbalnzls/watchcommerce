package app_context

import (
	"context"

	"go.uber.org/zap"

	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
	"github.com/iqbalnzls/watchcommerce/src/shared/logger"
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

func ParsingAppContext(ctx context.Context) *AppContext {
	appCtx, ok := ctx.Value(constant.AppContext).(*AppContext)
	if !ok {
		panic("please set app context")
	}

	return appCtx
}
