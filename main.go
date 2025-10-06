package main

import   (
	//"fmt",
	"log"
	"net/http"

	middlewares "github.com/aDiThYa-808/golang-http-server/internal/middlewares"
	handlers "github.com/aDiThYa-808/golang-http-server/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", middlewares.AuthMiddleware(http.HandlerFunc(handlers.HomeHandler)))
	mux.Handle("/stats",middlewares.AuthMiddleware(http.HandlerFunc(handlers.StatsHandler)))

	handler := middlewares.StatsRecorderMiddleware(mux)
	handler = middlewares.LoggingMiddleware(handler)
	handler = middlewares.RequestIdMiddleware(handler)

	log.Print("Server running on port 4000.")
	log.Fatal(http.ListenAndServe(":4000", handler))
}

