package internal

import (
	"context"

	"go.uber.org/zap"
)

var baseLogger *zap.Logger

func Init() error {
	l, err := zap.NewProduction()
	if err != nil {
		return err
	}
	baseLogger = l
	return nil
}

func Sync() error {
	if baseLogger != nil {
		return baseLogger.Sync()
	}
	return nil
}

type contextKey string

const RequestIDKey contextKey = "request_id"

func LoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	requestID, _ := ctx.Value(RequestIDKey).(string)

	if requestID == "" {
		return baseLogger.Sugar()
	}

	return baseLogger.
		With(zap.String("request_id", requestID)).
		Sugar()
}

func InforF(ctx context.Context, format string, args ...interface{}) {
	LoggerFromContext(ctx).Infof(format, args...)
}

func ErrorF(ctx context.Context, format string, args ...interface{}) {
	LoggerFromContext(ctx).Errorf(format, args...)
}
