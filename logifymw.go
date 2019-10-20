package logifymw
 
import (
    "log"
    "net/http"
    "time"
)
 
// LogIt logs the request method, path, and query as well as the elapsed time.
func LogIt(mux http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        defer measureTime(r.Method, r.URL.EscapedPath(), r.URL.RawQuery, time.Now())
        mux.ServeHTTP(w, r)
    }
}
 
func measureTime(method, path, query string, start time.Time) {
    elapsed := time.Since(start)
    log.Printf("%-4s %-50s %s", method, path + " " + query, elapsed)
}
