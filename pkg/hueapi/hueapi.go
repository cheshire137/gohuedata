package hueapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	ApiURL string
}

func NewClient(apiURL string) *Client {
	return &Client{ApiURL: apiURL}
}

func (a *Client) GetLights() ([]Light, error) {
	body, err := a.get("/lights")
	if err != nil {
		return nil, err
	}
	var lightResponse map[string]Light
	err = json.Unmarshal(body, &lightResponse)
	lights := []Light{}
	for _, light := range lightResponse {
		if light.UniqueID != "" {
			lights = append(lights, light)
		}
	}
	return lights, nil
}

func (a *Client) get(path string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, a.ApiURL+path, nil)
	if err != nil {
		return nil, err
	}
	return makeRequest(request)
}

func (a *Client) post(path string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, a.ApiURL+path, body)
	if err != nil {
		return nil, err
	}
	return makeRequest(request)
}

func makeRequest(request *http.Request) ([]byte, error) {
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
