package news // Kinda like exporting in js with module.export

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Article struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Urltoimage  string    `json:"urlToImage"`
	Publishedat time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type Results struct {
	Status		string		`json:"status"`
	TotalResult	int			`json:"totalResults"`
	Articles	[]Article	`json:"articles"`
}

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

func (c *Client) FetchEverything(query, page string) (*Results, error) {
	endpoint := fmt.Sprintf(
		"https://newsapi.org/v2/everything?q=%s&pageSize=%d&page=%s&apiKey=%s&sortBy=publishedAt&language=en", 
		url.QueryEscape(query), //URL encoding
		c.PageSize, 
		page, 
		c.key,
	)

	res, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Converts to byte slice 
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	
	// If response from News API is not ok, throw error
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	// If all is well, decode into Result struct using json.Unmarshall
	result := &Results{}
	return result, json.Unmarshal(body, result)
}