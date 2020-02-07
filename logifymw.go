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
		msg := fmt.Sprintf("%-15s %-4s %-50s %s %s", r.RemoteAddr, r.Method, r.URL.EscapedPath(), r.URL.RawQuery, r.UserAgent())
		defer measureTime(msg, time.Now())
		mux.ServeHTTP(w, r)
	}
}
func measureTime(msg string, start time.Time) {
	elapsed := time.Since(start)
	log.Printf("%s %s", msg, elapsed)
}
