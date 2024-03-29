package hue_api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	options "github.com/cheshire137/gohuedata/pkg/cli_options"
)

type Client struct {
	ApiURL     string
	Fahrenheit bool
}

func NewClient(apiURL string, fahrenheit bool) *Client {
	return &Client{ApiURL: apiURL, Fahrenheit: fahrenheit}
}

// https://developers.meethue.com/develop/hue-api/lights-api/#get-all-lights
func (c *Client) GetLights() (map[string]Light, error) {
	body, err := c.get("/lights")
	if err != nil {
		return nil, err
	}
	var lightResponse map[string]Light
	err = json.Unmarshal(body, &lightResponse)
	if err != nil {
		return nil, err
	}
	return lightResponse, nil
}

// https://developers.meethue.com/develop/hue-api/groupds-api/#get-group-attr
func (c *Client) GetGroup(id string) (*Group, error) {
	body, err := c.get(fmt.Sprintf("/groups/%s", id))
	if err != nil {
		return nil, err
	}
	var group Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

// https://developers.meethue.com/develop/hue-api/groupds-api/#get-all-groups
func (c *Client) GetGroups() (map[string]Group, error) {
	body, err := c.get("/groups")
	if err != nil {
		return nil, err
	}
	var groupResponse map[string]Group
	err = json.Unmarshal(body, &groupResponse)
	if err != nil {
		return nil, err
	}
	return groupResponse, nil
}

// https://developers.meethue.com/develop/hue-api/5-sensors-api/#get-all-sensors
func (c *Client) GetSensors(sensorSelection options.SensorSelection) ([]interface{}, error) {
	body, err := c.get("/sensors")
	if err != nil {
		return nil, err
	}
	var sensorResponse map[string]Sensor
	err = json.Unmarshal(body, &sensorResponse)
	if err != nil {
		return nil, err
	}
	sensors := make([]interface{}, 0, len(sensorResponse))
	loadTempSensors := sensorSelection != options.Motion
	loadMotionSensors := sensorSelection != options.Temperature
	for _, sensor := range sensorResponse {
		if sensor.IsTemperatureSensor() {
			if loadTempSensors {
				sensors = append(sensors, NewTemperatureSensor(sensor, c.Fahrenheit))
			}
		} else if sensor.IsMotionSensor() {
			if loadMotionSensors {
				sensors = append(sensors, NewMotionSensor(sensor))
			}
		} else {
			sensors = append(sensors, &sensor)
		}
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
