package news // Kinda like exporting in js with module.export

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Article Struct: Structure for each article within the result data
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

// FormatPublishedDate 
func (a *Article) FormatPublishedDate() string {
	year, month, day := a.Publishedat.Date()
	return fmt.Sprintf("%v %d, %d", month, day, year)  // Returns date formatted as month(string), day(integer) and year(integer)
}

// Results Struct for data received from the news api get request
type Results struct {
	Status		string		`json:"status"`
	TotalResults	int			`json:"totalResults"`
	Articles	[]Article	`json:"articles"`
}

// Client Struct reps the client
type Client struct {
	http	 *http.Client
	key		 string  	// API key
	PageSize int  		//Max number of results to return per page
}

func (c *Client) FetchTopHeadlines(page string) (*Results, error) {
	endpoint := fmt.Sprintf(
		"https://newsapi.org/v2/top-headlines?country=jp&pageSize=%d&page=%s&apiKey=%s&sortBy=publishedAt&language=jp",
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

// NewClient Creates and returns a new Client
func NewClient(httpClient *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}