package news // Kinda like exporting in js with module.export

import "net/http"

// Client Struct reps the client
type Client struct {
	http	 *http.Client
	key		 string  	// API key
	PageSize int  		//Max number of results to return per page
}

// Creates and returns a new Client
func NewClient(httpClient *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}