package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type key int

const requestIDKey key = iota

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get("X-Request-ID")

		if rid == "" {
			rid = uuid.New().String()
		}

		w.Header().Set("X-Request-ID", rid)

		ctx := context.WithValue(r.Context(), requestIDKey, rid)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestID(ctx context.Context) string {
	v := ctx.Value(requestIDKey)

	if v == nil {
		return ""
	}

	s, ok := v.(string)

	if ok {
		return s
	}

	return ""
}
