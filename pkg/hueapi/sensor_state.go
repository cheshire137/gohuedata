package hueapi

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

func (s *SensorState) FahrenheitTemperature() int {
	celsiusTemp := s.Temperature / 100 // API returns a value like 2277 for 22.77°C
	return (celsiusTemp * 9 / 5) + 32
}
