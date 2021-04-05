package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http" // Provides HTTP client and server implementation for use in the app
	"net/url"
	"os" // Allows operating system functionalities

	"github.com/joho/godotenv"
)

 var tpl = template.Must(template.ParseFiles("index.html"))  // Package level variable. Parses the index.html file as a template

// w represents the "res" and r is "req"
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

// function to deal with users' searches
func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	searchQuery := params.Get("q")  //Users' query
	page := params.Get("page")  // Page through results
	if page == "" {
		page = "1"
	}

	fmt.Println("Search Query is: ", searchQuery)
	fmt.Println("Page is: ", page)
}

// This is like the renderer for the main page
func main() {
	err := godotenv.Load()  // Reads the .env file
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")  // Sets the port
	if port == "" {
		port = "5000"
	}

	fs := http.FileServer(http.Dir("assets"))  // Pass the directory where all static files are saved to the file server

	mux := http.NewServeMux()  // http request multiplexer (need to read up on it more)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))  // StripPrefix strips off the specified prefix
	mux.HandleFunc("/search", searchHandler)
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)  // starts the server on port
}


/* 
	1. Run -> go mod init `github.com/username/repo` to initialize module
	2. Run -> go build
*/