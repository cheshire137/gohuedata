package hueapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	ApiURL   string
	Username string
}

func NewClient(apiURL string, username string) *Client {
	return &Client{ApiURL: apiURL, Username: username}
}

func (a *Client) GetLights() ([]Light, error) {
	body, err := a.get("/lights")
	if err != nil {
		return nil, err
	}
	var lightResponse map[string]Light
	err = json.Unmarshal(body, &lightResponse)
	lights := make([]Light, len(lightResponse))
	for _, light := range lightResponse {
		lights = append(lights, light)
	}
	return lights, nil
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
	return fmt.Sprintf("%s/%s%s", a.ApiURL, a.Username, path)
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
