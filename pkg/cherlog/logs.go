package cherlog

import (
	"context"
	"go.uber.org/zap"
	"sync"
)

var (
	once sync.Once
	log  *zap.Logger
)

type Key string

func (k Key) String() string {
	return string(k)
}

const (
	TraceID       Key = "trace_id"
	CorrelationID Key = "correlation_id"
	RequestID     Key = "request_id"
)

func NewLog() *zap.Logger {
	once.Do(func() {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		log = logger
	})

	return log
}

type LoggerStruct struct{}

var LoggerIDKey = LoggerStruct{}

func SetLogFromCtx(ctx context.Context, logger *zap.Logger) context.Context {
	if logger == nil {
		logger = NewLog()
	}

	return context.WithValue(ctx, LoggerIDKey, logger)
}

func GetLogFromCtx(ctx context.Context) *zap.Logger {
	if log, ok := ctx.Value(LoggerIDKey).(*zap.Logger); ok {
		return log
	}
	return NewLog()
}
