package middlewares

import "net/http"

type LoggingResponseWriter struct {
	http.ResponseWriter
	Status int
	Body   []byte
}

func (lrw *LoggingResponseWriter) WriteHeader(statusCode int) {
	lrw.Status = statusCode
	lrw.ResponseWriter.WriteHeader(statusCode)
}

func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	lrw.Body = append(lrw.Body, b...)
	return lrw.ResponseWriter.Write(b)
}
