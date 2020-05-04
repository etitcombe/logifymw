# logifymw

A silly little logging middleware.

It has three methods, LogIt, LogItMore, and LogItMoreMore, which all take an
[http.Handler](https://golang.org/pkg/net/http/#Handler) and return an 
[http.Handler](https://golang.org/pkg/net/http/#Handler).

LogIt logs the method, path, and query as well as the elapsed time of each request.

LogItMore logs the remote address, method, path, query, and user agent as well
as the elapsed time of each request.

LogItMoreMore adds the status and the size of the response to what was already
added by LogItMore.

Use it to wrap a mux so that it applies to all the routes:

    srv := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: logifymw.LogIt(mux)}

Or, just to individual HandlerFuncs:

    http.HandleFunc("/some/path/", logifymw.LogIt(somePathHandler()))
