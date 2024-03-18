package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		startTimeStr := startTime.Format("2006-01-02 15:04:05.000-0700")

		next.ServeHTTP(w, r)

		endTimeStr := time.Now().Format("2006-01-02 15:04:05.000-0700")

		logrus.WithFields(logrus.Fields{
			"start":      startTimeStr,
			"end":        endTimeStr,
			"durationMs": time.Since(startTime).Milliseconds(),
			"remoteAddr": r.RemoteAddr,
			"method":     r.Method,
			"url":        r.URL,
			"proto":      r.Proto,
			"cid":        r.URL.Query().Get("companyId"),
			"uid":        r.URL.Query().Get("userId"),
			"permission": r.URL.Query().Get("permissions"),
			"traceId":    r.Header.Get("X-Daji-Trace-Id"),
		}).Info("Request Fulfilled")
	})
}
