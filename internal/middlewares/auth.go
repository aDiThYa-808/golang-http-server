package middlewares;

import (
	"net/http";
)

func AuthMiddleware(next http.Handler) http.Handler{
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

