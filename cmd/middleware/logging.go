package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type (
	Logger struct {
		handler http.Handler
		slog    *slog.Logger
	}

	loggingResponseWriter struct {
		w          http.ResponseWriter
		statusCode int
		data       []byte
	}
)

// Logs and prettifies the request, along with the error
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	lrw := NewLoggingResponseWriter(w)
	l.handler.ServeHTTP(lrw, r)

	l.slog.Info("route",
		"method", r.Method,
		"path", r.URL.Path,
		"elasped", time.Since(start),
		// "data", string(lrw.data),
		"status", fmt.Sprintf("%d %s", lrw.statusCode, http.StatusText(lrw.statusCode)),
	)
}

func (l *loggingResponseWriter) Header() http.Header {
	return l.w.Header()
}

func (l *loggingResponseWriter) Write(data []byte) (int, error) {
	l.data = append(l.data, data...)
	return l.w.Write(data)
}

func (l *loggingResponseWriter) WriteHeader(code int) {
	l.statusCode = code
	l.w.WriteHeader(code)
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler, slog *slog.Logger) *Logger {
	return &Logger{
		handler: handlerToWrap,
		slog:    slog,
	}
}

// NewLoggingResponseWriter contructs a wrapper for ResponseWriter,
// so we can read values not easily exposed
func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{
		w:          w,
		statusCode: http.StatusOK,
		data:       []byte{},
	}
}
