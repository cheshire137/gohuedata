package hueapi

type SensorEvent struct {
	ButtonEvent int    `json:"buttonevent"`
	EventType   string `json:"eventtype"`
}
