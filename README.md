# logifymw

A silly little logging middleware.

It has one method, LogIt, which takes an [http.Handler](https://golang.org/pkg/net/http/#Handler) and returns an [http.HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc). It logs the method, path, and query as well as the elapsed time of each request.

Use it to wrap a mux so that it applies to all the routes:

    srv := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: logifymw.LogIt(mux)}

Or, just to individual HandlerFuncs:

    http.HandleFunc("/some/path/", logifymw.LogIt(somePathHandler()))
