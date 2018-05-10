package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("REQUEST:", r)
	return
}

// Logger Middleware
func logginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do stuff 
		fmt.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Print to test
	fmt.Println("Hello, these are the humble beginnings!")

	// create a new router
	r := mux.NewRouter()

	// index route
	r.HandleFunc("/", HomeHandler)

	// register logginMiddleware
	r.Use(logginMiddleware)

	// register router with http
	http.Handle("/", r)

	// start serve
	http.ListenAndServe(":8000", nil)
}

