package main;

import (
	//"fmt",
	"net/http";
	"log";
	middlewares "github.com/aDiThYa-808/golang-http-server/internal/middlewares"
)

func main(){
	mux:= http.NewServeMux();
	mux.Handle("/",middlewares.AuthMiddleware(http.HandlerFunc(homeHandler)));

	log.Print("Server running on port 4000.");
	log.Fatal(http.ListenAndServe(":4000",middlewares.LoggingMiddleware(mux)));
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Home - Welcome back!!\n"));
}
