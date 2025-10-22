package middlewares

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedUser := "admin"
		expectedPassword := "pass" // I have hardcoded the password because this is just a learning/personal project and i want to focus on the other important features.

		user, pass, ok := r.BasicAuth()

		if !ok || user != expectedUser || pass != expectedPassword {
			w.Header().Set("WWW-Authenticate", `Basic realm="protected"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
