package logifymw

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// LogIt logs the request method, path, and query as well as the elapsed time.
func LogIt(mux http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%-4s %-50s", r.Method, r.URL.EscapedPath()+" "+r.URL.RawQuery)
		defer measureTime(msg, time.Now())
		mux.ServeHTTP(w, r)
	}
}

// LogItMore logs the request address, method, path, query, and user agent as well
// as the elapsed time.
func LogItMore(mux http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%-15s %-4s %-50s %s", r.RemoteAddr, r.Method, r.URL.EscapedPath()+" "+r.URL.RawQuery, r.UserAgent())
		defer measureTime(msg, time.Now())
		mux.ServeHTTP(w, r)
	}
}

// LogItMoreMore adds the status and the size to the log message to what was
// already added by LogItMore
func LogItMoreMore(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		lw := loggingResponseWriter{w, 0, 0}
		h.ServeHTTP(&lw, r)

		msg := fmt.Sprintf("%-15s %-4s %-50s %s %d %d", r.RemoteAddr, r.Method, r.URL.EscapedPath()+" "+r.URL.RawQuery, r.UserAgent(), lw.status, lw.size)
		log.Printf("%s %s", msg, time.Since(now))
	})
}

func measureTime(msg string, start time.Time) {
	elapsed := time.Since(start)
	log.Printf("%s %s", msg, elapsed)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *loggingResponseWriter) Write(data []byte) (int, error) {
	written, err := w.ResponseWriter.Write(data)
	w.size += written
	return written, err
}

func (w *loggingResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.status = statusCode
}
