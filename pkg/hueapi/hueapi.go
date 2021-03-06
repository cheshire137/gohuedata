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

// https://developers.meethue.com/develop/hue-api/lights-api/#get-all-lights
func (c *Client) GetLights() ([]Light, error) {
	body, err := c.get("/lights")
	if err != nil {
		return nil, err
	}
	var lightResponse map[string]Light
	err = json.Unmarshal(body, &lightResponse)
	if err != nil {
		return nil, err
	}
	lights := []Light{}
	for _, light := range lightResponse {
		if light.UniqueID != "" {
			lights = append(lights, light)
		}
	}
	return lights, nil
}

// https://developers.meethue.com/develop/hue-api/5-sensors-api/#get-all-sensors
func (c *Client) GetSensors() ([]Sensor, error) {
	body, err := c.get("/sensors")
	if err != nil {
		return nil, err
	}
	var sensorResponse map[string]Sensor
	err = json.Unmarshal(body, &sensorResponse)
	if err != nil {
		return nil, err
	}
	sensors := make([]Sensor, 0, len(sensorResponse))
	for _, sensor := range sensorResponse {
		sensors = append(sensors, sensor)
	}
	return sensors, nil
}

func (c *Client) get(path string) ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, c.ApiURL+path, nil)
	if err != nil {
		return nil, err
	}
	return makeRequest(request)
}

func (c *Client) post(path string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, c.ApiURL+path, body)
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
