package av

import (
	"io"
    "fmt"
	"log"
	"net/http"
)

var ()

type client struct {
	api_key string
	host    string
}

func NewClient() *client {
	return &client{
		api_key: "C227WD9W3LUVKVV9",
		host:    "https://www.alphavantage.co/",
	}

}

func (c *client) Get() {
	req, err := http.NewRequest(http.MethodGet, c.host, nil)
	if err != nil {
		log.Printf("Failed to construct request\n")
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {

		log.Printf("Error from server %s\n", err)
	}

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Printf("Could not read the response body")
	}

	fmt.Printf("Response: %s", resBody)
}
