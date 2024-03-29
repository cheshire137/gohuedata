package hue_api

import (
	"fmt"
	"strings"
	"time"
)

const LastUpdatedFormat = "2006-01-02T15:04:05"

type SensorState struct {
	LightLevel  int    `json:"lightlevel"`
	Dark        bool   `json:"dark"`
	Daylight    bool   `json:"daylight"`
	LastUpdated string `json:"lastupdated"`
	ButtonEvent int    `json:"buttonevent"`
	Flag        bool   `json:"flag"`
	Status      int    `json:"status"`
	Temperature int    `json:"temperature"`
}

func (s *SensorState) LastUpdatedAt() (*time.Time, error) {
	if s.LastUpdated == "" {
		return nil, fmt.Errorf("no last update time")
	}
	lastUpdatedTime, err := time.Parse(LastUpdatedFormat, s.LastUpdated)
	if err != nil {
		return nil, err
	}
	_, timezoneOffset := time.Now().Zone()
	lastUpdatedTime = lastUpdatedTime.Add(time.Duration(timezoneOffset) * time.Second)
	return &lastUpdatedTime, nil
}

func (s *SensorState) LastUpdatedSummary() string {
	lastUpdated, err := s.LastUpdatedAt()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s on %s", lastUpdated.Format("3:04 PM"), lastUpdated.Format("Jan 2, 2006"))
}

func (s *SensorState) CelsiusTemperature() float32 {
	return float32(s.Temperature) / 100 // API returns a value like 2277 for 22.77°C
}

func CelsiusToFahrenheit(celsius float32) float32 {
	return (celsius * 9 / 5) + 32
}

func (s *SensorState) FahrenheitTemperature() float32 {
	return CelsiusToFahrenheit(s.CelsiusTemperature())
}

func FahrenheitToCelsius(fahrenheit float32) float32 {
	return (fahrenheit - 32) * 5 / 9
}

func ConvertTemperature(temperature float32, units string, fahrenheit bool) float32 {
	units = strings.ToUpper(units)
	if fahrenheit && units == "C" {
		return CelsiusToFahrenheit(temperature)
	}
	if !fahrenheit && units == "F" {
		return FahrenheitToCelsius(temperature)
	}
	return temperature
}
