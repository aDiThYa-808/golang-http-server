package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v0 := time.Now()

		rid := GetRequestID(r.Context())

		log.Println("Request method: ", r.Method)
		log.Println("Request path: ", r.URL.Path)
		log.Println("Request ID: ", rid)
		log.Println("Request timestamp: ", v0.Format(time.RFC3339))

		next.ServeHTTP(w, r)
		v1 := time.Since(v0)
		log.Println("Time taken to complete request: ", v1)
	})
}
