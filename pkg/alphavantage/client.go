package av

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
    "crypto/tls"
)

var ()

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	Api_key string
	Host    string
	Client  HTTPClient
}

type Result struct {
	Meta interface{}    `json:"Meta Data,omitempty"`
	Data map[string]Day `json:"Time Series (Daily),omitempty"`
}

type Day struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func NewClient() *Client {
	return &Client{
		Api_key: "C227WD9W3LUVKVV9",
		Host:    "https://www.alphavantage.co/",
        // Skipping validation of TLS certificates for now
        // allows use of scratch images. Would need to copy the CA root into
        // scratch image
		Client: &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}},
	}

}

func (c *Client) Get() (*Result, error) {
	return c.GetSymbol("MSFT")
}

func (c *Client) GetSymbol(Symbol string) (*Result, error) {
	//query?apikey=C227WD9W3LUVKVV9&function=TIME_SERIES_DAILY&symbol=MSFT
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/query?apikey=%s&function=TIME_SERIES_DAILY&symbol=%s", c.Host, c.Api_key, Symbol), nil)
	if err != nil {
		log.Printf("Failed to construct request\n")
		return nil, err
	}

	res, err := c.Client.Do(req)

	if err != nil {
		log.Printf("Error from server %s\n", err)
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	result := Result{}
	err = json.NewDecoder(bytes.NewReader(resBody)).Decode(&result)

	return &result, err
}
