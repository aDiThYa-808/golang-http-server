package main

import   (
	//"fmt",
	"log"
	"net/http"

	middlewares "github.com/aDiThYa-808/golang-http-server/internal/middlewares"
	homeHandler "github.com/aDiThYa-808/golang-http-server/handlers/home"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", middlewares.AuthMiddleware(http.HandlerFunc(homeHandler.HomeHandler)))

	handler := middlewares.StatsRecorderMiddleware(mux)
	handler = middlewares.LoggingMiddleware(handler)
	handler = middlewares.RequestIdMiddleware(handler)

	log.Print("Server running on port 4000.")
	log.Fatal(http.ListenAndServe(":4000", handler))
}

