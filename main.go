package main

import (
	"fmt"
	"net/http"
)

func main() {
	// define the server. here & represents that we storing the value as a pointer by taking the
	// memory address (reference type) rather than the value (value type)
	server := &http.Server{
		// inside here we can define our server properties like
		Addr:    ":3000", // port
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
