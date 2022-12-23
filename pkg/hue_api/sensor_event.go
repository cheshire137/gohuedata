package hue_api

type SensorEvent struct {
	ButtonEvent int    `json:"buttonevent"`
	EventType   string `json:"eventtype"`
}
