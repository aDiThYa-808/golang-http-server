package middlewares

import (
	"log"
	"net/http"
)

func MaxBodySize(maxBytes int64) func(http.Handler) http.Handler{
	return func(next http.Handler) http.Handler{
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			log.Println("MaxBodySize middleware")
			r.Body = http.MaxBytesReader(w,r.Body,maxBytes)
			next.ServeHTTP(w,r)
		})
	}
}
