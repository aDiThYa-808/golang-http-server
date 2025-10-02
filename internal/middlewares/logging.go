package middlewares;

import (
	"net/http";
	"log";
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		log.Println("Request method: ",r.Method);
		log.Println("Request path: ",r.URL.Path);
		log.Println("Request timestamp: ",time.Now().Format(time.RFC3339));

		v0 := time.Now();
		next.ServeHTTP(w,r);
		v1 := time.Since(v0);
		log.Println("Time taken to complete request: ",v1);
	});
}
