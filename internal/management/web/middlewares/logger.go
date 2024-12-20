package middlewares

import (
	"bff/pkg"
	"bff/pkg/cherlog"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

func WithRequestLogger(ctx context.Context, request *http.Request) context.Context {
	correlationIDStr := request.Header.Get("X-Correlation-ID")
	correlationID, err := uuid.Parse(correlationIDStr)
	if err != nil {
		correlationID = pkg.GetCorrelationID(ctx)
	}

	requestIDStr := request.Header.Get("X-Request-ID")
	requestID, err := uuid.Parse(requestIDStr)
	if err != nil {
		requestID = pkg.GetCorrelationID(ctx)
	}

	var body []byte
	if request.Body != nil {
		defer request.Body.Close()
		request.Body.Read(body)
	}

	logger := cherlog.GetLogFromCtx(ctx).With(
		zap.String("http.request.method", request.Method),
		zap.String("http.request.url", request.URL.Path),
		zap.String("http.request.agent", request.UserAgent()),
		zap.String("http.request.correlation_id", correlationID.String()),
		zap.String("http.request.request_id", requestID.String()),
		zap.String("http.request.remote_addr", request.RemoteAddr),
		zap.String("http.request.body", string(body)),
	)

	return cherlog.SetLogFromCtx(ctx, logger)
}
