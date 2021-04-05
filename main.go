package main

import (
	"net/http"  // Provides HTTP client and server implementation for use in the app
	"os" // Allows operating system functionalities
)

// w represents the "res" and r is "req"
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	mux := http.NewServeMux()  // http request multiplexer (need to read up on it more)
	
	mux.HandleFunc("/", indexHandler)
	
	http.ListenAndServe(":"+port, mux)  // starts the server on port
}


/* 
	1. Run -> go mod init `github.com/username/repo` to initialize module
	2. Run -> go build
*/