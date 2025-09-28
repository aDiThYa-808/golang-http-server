package main;

import (
	//"fmt",
	"net/http";
	"time";
	"log"
)

func main(){
	mux:= http.NewServeMux();
	mux.Handle("/",authMiddleware(http.HandlerFunc(homeHandler)));
	mux.HandleFunc("/time",timeHandler);

	
	log.Print("Server running on port 4000.");
	log.Fatal(http.ListenAndServe(":4000",loggingMiddleware(mux)));
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Home - Welcome back!!\n"));
}

func timeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/plain");
	w.Write([]byte(time.Now().Format(time.RFC3339+"\n")));
} 

func authMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		expectedUser:= "adi";
		expectedPassword:= "sybau";

		user,pass,ok := r.BasicAuth();

		if !ok || user != expectedUser || pass != expectedPassword{
			w.Header().Set("WWW-Authenticate",`Basic realm="protected"`);
			http.Error(w,"Unauthorized",http.StatusUnauthorized);
			return;
		}

		next.ServeHTTP(w,r);
	});
}

func loggingMiddleware(next http.Handler) http.Handler{
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
