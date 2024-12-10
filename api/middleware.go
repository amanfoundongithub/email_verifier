package api

import (
	"log"
	"net/http"
	"time"
)

func addMiddleware(next http.Handler) http.Handler {

	// Returns an HTTP Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Log the input time
		start := time.Now()

		// Log the arrival
		log.Printf("%s REQUEST RECEIVED ON %s", r.Method, r.URL.Path)

		// Next handler
		next.ServeHTTP(w, r)

		// Output
		log.Printf("Service Completed In : %v", time.Since(start))
	})

}

func addMiddlewareToMUX(mux *http.ServeMux) http.Handler {

	// Wrap the middleware around
	return addMiddleware(mux)
}