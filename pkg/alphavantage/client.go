package av

import (
    "bytes"
	"fmt"
	"io"
	"log"
	"net/http"
    "encoding/json"
)

var ()

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	api_key string
	host    string
	client  HTTPClient
}

type Result struct {
    Meta interface{} `json:"Meta Data,omitempty"`
    Data map[string]Day `json:"Time Series (Daily),omitempty"`
}

type Day struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func NewClient() *client {
	return &client{
		api_key: "C227WD9W3LUVKVV9",
		host:    "https://www.alphavantage.co/",
		client:  http.DefaultClient,
	}

}

func (c *client) Get() (*Result, error) {
	//query?apikey=C227WD9W3LUVKVV9&function=TIME_SERIES_DAILY&symbol=MSFT
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/query?apikey=%s&function=TIME_SERIES_DAILY&symbol=%s", c.host, c.api_key, "MSFT"), nil)
	if err != nil {
		log.Printf("Failed to construct request\n")
		return nil, err
	}

	res, err := c.client.Do(req)

	if err != nil {
		log.Printf("Error from server %s\n", err)
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
    result := Result{}
    err = json.NewDecoder(bytes.NewReader(resBody)).Decode(&result)

	return &result, err
}
