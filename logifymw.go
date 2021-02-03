package logifymw

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// LogIt logs the request method, path, and query as well as the elapsed time.
func LogIt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%-4s %-50s", r.Method, r.URL.Path+" "+unescape(r.URL.RawQuery))
		defer measureTime(msg, time.Now())
		next.ServeHTTP(w, r)
	})
}

// LogIt2 logs to log the request method, path and query, the response status, size, and the elapsed time.
func LogIt2(log *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		lw := loggingResponseWriter{w, http.StatusOK, 0}
		next.ServeHTTP(&lw, r)

		msg := fmt.Sprintf("%-8s %-71s %d %-5d", r.Method, unescape(r.URL.RequestURI()), lw.status, lw.size)
		log.Printf("%s %s", msg, time.Since(now))
	})
}

// LogItMore logs the request address, method, path, query, and user agent as well
// as the elapsed time.
func LogItMore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("%-15s %-4s %-50s %s", getIP(r), r.Method, r.URL.Path+" "+unescape(r.URL.RawQuery), r.UserAgent())
		defer measureTime(msg, time.Now())
		next.ServeHTTP(w, r)
	})
}

// LogItMoreMore adds the status and the size to the log message to what was
// already added by LogItMore.
func LogItMoreMore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		lw := loggingResponseWriter{w, http.StatusOK, 0}
		next.ServeHTTP(&lw, r)

		msg := fmt.Sprintf("%-15s %-4s %-50s %s %d %d", getIP(r), r.Method, r.URL.Path+" "+unescape(r.URL.RawQuery), r.UserAgent(), lw.status, lw.size)
		log.Printf("%s %s", msg, time.Since(now))
	})
}

// LogItMoreMore2 adds the protocol, and uses the request URI instead of the path
// and query, to the log message to what was already added by LogItMoreMore.
func LogItMoreMore2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		lw := loggingResponseWriter{w, http.StatusOK, 0}
		next.ServeHTTP(&lw, r)

		msg := fmt.Sprintf("%-15s %-8s %-4s %-50s %s %d %d", getIP(r), r.Proto, r.Method, unescape(r.URL.RequestURI()), r.UserAgent(), lw.status, lw.size)
		log.Printf("%s %s", msg, time.Since(now))
	})
}

func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

func unescape(input string) string {
	output, err := url.QueryUnescape(input)
	if err != nil {
		log.Printf("warning: unable to unescape the input: %v\n", err)
		output = input
	}
	return output
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
