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
		log.Println("Total Requests: ",GetStats("total"))
		log.Println("Successful Requests: ",GetStats("success"))
		log.Println("Client errors: ",GetStats("client_error"))
		log.Println("unauthorized Requests: ",GetStats("unauthorized"))
		log.Println("Server errors: ",GetStats("server_error"))




		next.ServeHTTP(w, r)
		v1 := time.Since(v0)
		log.Println("Time taken to complete request: ", v1)
	})
}
