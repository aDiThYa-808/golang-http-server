# golang-http-server

A minimal HTTP server implemented using Go’s net/http package. Implemented to understand how handlers, middleware, concurrency, and context propagation work in Go.

## Features (in the current stage)
### Endpoints:
- **`/`**: endpoint that is protected by an authentication middleware.
- **`/stats`**: endpoint that returns the request metrics like total requests, successful requests and Unauthorized requests in JSON.

### Middlewares:
- **`AuthMiddleware()`**: protects the endpoints and returns 401 if the Basic Auth credentials are invalid.
- **`LoggingMiddleware()`**: logs the request's method, path, timestamp and time taken to complete the request.
- **`StatsMiddleware()`**: the outermost middleware that checks the request's status code and updates the in-memory counters using `atomic.AddInt64()` to keep track of the metrics.

## Go Packages used
- `net/http`
- `context`
- `sync/atomic`
- `encoding/json`
- `time`
- `log`

---

# Important Note !
**This project is still under development. The README will be updated with new features and endpoints in the coming days.**

