package hueapi

type SensorState struct {
	LightLevel  int    `json:"lightlevel"`
	Dark        bool   `json:"dark"`
	Daylight    bool   `json:"daylight"`
	LastUpdated string `json:"lastupdated"`
	ButtonEvent int    `json:"buttonevent"`
	Flag        bool   `json:"flag"`
	Status      int    `json:"status"`
}
