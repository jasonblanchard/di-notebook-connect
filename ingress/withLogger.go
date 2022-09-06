package ingress

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type ResponseRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *ResponseRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

// Logger is a middleware handler that does request logging
type WithLogger struct {
	handler http.Handler
	logger  *zap.SugaredLogger
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *WithLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseRecorder := &ResponseRecorder{
		ResponseWriter: w,
		Status:         http.StatusOK,
	}

	start := time.Now()

	l.handler.ServeHTTP(responseRecorder, r)

	l.logger.Infow("ingress",
		"method", r.Method,
		"path", r.URL.Path,
		"host", r.Host,
		"protocol", r.Proto,
		"status", responseRecorder.Status,
		"x-forwarded-for", r.Header.Get("x-forwarded-for"),
		"x-forwarded-proto", r.Header.Get("x-forwarded-proto"),
		"durationS", time.Since(start),
	)
}

// NewLogger constructs a new Logger middleware handler
func NewWithLogger(handlerToWrap http.Handler, logger *zap.SugaredLogger) *WithLogger {
	return &WithLogger{
		handler: handlerToWrap,
		logger:  logger,
	}
}
