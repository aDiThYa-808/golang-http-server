package middlewares

import (
	"sync/atomic"
	"net/http"
)

var totalRequests int64;
var successfulRequests int64;
var serverError int64;
var clientError int64;
var unauthorizedRequests int64;

type statusResponseWriter struct{
	http.ResponseWriter
	status int
}

func (w *statusResponseWriter) WriteHeader(code int){
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func StatsRecorderMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		newResponseWriter := &statusResponseWriter{ResponseWriter:w,status:200}

		next.ServeHTTP(newResponseWriter,r)

		status := newResponseWriter.status

		incrementCounter(status)

	})
}

func incrementCounter(status int){

	atomic.AddInt64(&totalRequests,1)

	if status >= 200 && status < 300{
		atomic.AddInt64(&successfulRequests,1)
	}
	if status == 401{
		atomic.AddInt64(&unauthorizedRequests,1)
	}
	if status >= 400 && status <500{
		atomic.AddInt64(&clientError,1)
	}
	if status >=500 && status < 600{
		atomic.AddInt64(&serverError,1)
	}
}

func GetStats(statType string) int64 {
	switch statType{
	case "total":
		return atomic.LoadInt64(&totalRequests)
	case "success":
		return atomic.LoadInt64(&successfulRequests)
	case "server_error":
		return atomic.LoadInt64(&serverError)
	case "client_error":
		return atomic.LoadInt64(&clientError)
	case "unauthorized":
		return atomic.LoadInt64(&unauthorizedRequests)
	default:
		return 0
	}

}




