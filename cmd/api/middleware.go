package api

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %v %v %v\n", r.Method, r.RemoteAddr, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
