package main

import (
	//"fmt",
	"log"
	"net/http"

	middlewares "github.com/aDiThYa-808/golang-http-server/internal/middlewares"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", middlewares.AuthMiddleware(http.HandlerFunc(homeHandler)))

	handler := middlewares.LoggingMiddleware(mux)
	handler = middlewares.RequestIdMiddleware(handler)

	log.Print("Server running on port 4000.")
	log.Fatal(http.ListenAndServe(":4000", handler))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home - Welcome back!!\n"))
}
