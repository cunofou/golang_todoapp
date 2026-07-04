package core_http_middleware

import (
	"context"
	"net/http"

	core_logger "github.com/cunofou/golang_todoapp/internal/core/logger"
	core_http_response "github.com/cunofou/golang_todoapp/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = uuid.NewString()
			}
			r.Header.Set("X-Request-Id", requestID)
			w.Header().Set("X-Request-Id", requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)
			ctx := context.WithValue(r.Context(), "log", l)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(p, "during handle http got panic")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}

	func Trace() Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := r.Context()
            log := core_logger.FromContext(ctx)
            rw := core_http_response.NewHTTPResponseWriter(w)

            before := time.Now()
            log.Debug(
                ">>> incoming HTTP request",
                zap.Time("time", before), // или before.UTC()
            )
            next.ServeHTTP(rw, r)

            log.Debug(
                "<<< done HTTP response",
                zap.Int("status-code", rw.GetStatusCodeOrPanic()),
                zap.Duration("latency", time.Since(before)),
            )
        })
    }
}