package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func HttpLogger(logger *slog.Logger) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrapper := &WrapperWriter{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
			}
			next.ServeHTTP(wrapper, r)
			end := time.Since(start)
			logger.Debug(r.URL.String(),
				slog.String("Method", r.Method),
				slog.Int("Status", wrapper.StatusCode),
				slog.String("Time", end.String()))
		})
	}
}
