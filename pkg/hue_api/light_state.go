package hue_api

type LightState struct {
	On               bool      `json:"on"`
	Alert            string    `json:"alert"`
	ColorMode        string    `json:"colormode"`
	Mode             string    `json:"mode"`
	Reachable        bool      `json:"reachable"`
	Brightness       int       `json:"bri"`
	Hue              int       `json:"hue"`
	Saturation       int       `json:"sat"`
	Effect           string    `json:"effect"`
	XY               []float32 `json:"xy"`
	ColorTemperature int       `json:"ct"`
}
