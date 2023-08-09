package av

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

type testClient struct {
}

func (t *testClient) Do(req *http.Request) (*http.Response, error) {
	file, _ := os.Open("stock.json")
    defer file.Close()
	json, _ := io.ReadAll(file) 
	r := io.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{StatusCode: http.StatusOK, Body: r}, nil

}

func TestTestClientGet(t *testing.T) {
	client := NewClient()
	client.Client = &testClient{}

	client.Get()
}

func TestGetRealData(t *testing.T) {
    client := NewClient()

    resp, err := client.Get()
    if err != nil {
        t.Fatal(err)
    }

    fmt.Printf("%v", resp)
}
