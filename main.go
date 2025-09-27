package main;

import (
	//"fmt",
	"net/http";
	"time";
	"log"
)

func main(){
	mux:= http.NewServeMux();

	mux.HandleFunc("/ping",pingHandler);
	mux.HandleFunc("/time",timeHandler);

	log.Fatal(http.ListenAndServe(":4000",mux));
}

func pingHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/plain")
	w.Write([]byte("pong\n"));
}

func timeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/plain");
	w.Write([]byte(time.Now().Format(time.RFC3339+"\n")));
}
