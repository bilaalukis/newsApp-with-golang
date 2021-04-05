package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http" // Provides HTTP client and server implementation for use in the app
	"net/url"
	"os" // Allows operating system functionalities
	"time"

	"github.com/bilaalukis/newsApp-with-golang.git/news"
	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))  // Package level variable. Parses the index.html file as a template


// w represents the "res" and r is "req"
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

// function to deal with users' searches
func searchHandler(newsapi *news.Client) http.HandlerFunc {
	// using closure to access the newsapi client
	return func(w http.ResponseWriter, r *http.Request) {
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

		results, err := newsapi.FetchEverything(searchQuery, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
	
		fmt.Println("%+v", results)
	} 
}

// This is like the renderer for the main page
func main() {
	err := godotenv.Load()  // Reads the .env file
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")  // Sets the port from .env
	if port == "" {
		port = "5000"
	}

	apiKey := os.Getenv("NEWS_API_KEY")  // Sets the api key from .env
	if apiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	myClient := &http.Client{Timeout: 10 * time.Second}  // Timeout for requested resources to be sent. Should not be more than 10 seconds.
	newsapi := news.NewClient(myClient, apiKey, 20)

	fs := http.FileServer(http.Dir("assets"))  // Pass the directory where all static files are saved to the file server

	mux := http.NewServeMux()  // http request multiplexer (need to read up on it more)
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))  // StripPrefix strips off the specified prefix
	mux.HandleFunc("/search", searchHandler(newsapi))
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)  // starts the server on port
}


/* 
	1. Run -> go mod init `github.com/username/repo` to initialize module
	2. Run -> go build
*/