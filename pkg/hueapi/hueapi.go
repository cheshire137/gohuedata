package hueapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL      string
	Username string
}

func NewClient(url string, username string) *Client {
	return &Client{URL: url, Username: username}
}

func (a *Client) get(path string) ([]byte, error) {
	fmt.Println(a.urlFor(path))
	request, err := http.NewRequest(http.MethodGet, a.urlFor(path), nil)
	if err != nil {
		return nil, err
	}
	return a.makeRequest(request)
}

func (a *Client) post(path string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, a.urlFor(path), body)
	if err != nil {
		return nil, err
	}
	return a.makeRequest(request)
}

func (a *Client) urlFor(path string) string {
	return fmt.Sprintf("%s/api/%s%s", a.URL, a.Username, path)
}

func (a *Client) makeRequest(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d %s", response.StatusCode, response.Status)
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
