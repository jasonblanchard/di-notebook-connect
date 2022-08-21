package ingress

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Logger is a middleware handler that does request logging
type WithLogger struct {
	handler http.Handler
	logger  *zap.SugaredLogger
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *WithLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	l.logger.Infow("handled",
		"method", r.Method,
		"path", r.URL.Path,
		"host", r.Host,
		"protocol", r.Proto,
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
