package middleware

import (
	"log"
	"net/http"
)

func HttpLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Log: Connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func HttpTracer(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Trace: Request for %s", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
