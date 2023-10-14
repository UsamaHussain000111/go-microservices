package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	//add a chi router
	router := chi.NewRouter()
	//use the logging for routes
	router.Use(middleware.Logger)

	router.Get("/hello", basicHandler)
	// define the server. here & represents that we storing the value as a pointer by taking the
	// memory address (reference type) rather than the value (value type)
	server := &http.Server{
		// inside here we can define our server properties like
		Addr:    ":3000", // port
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
