package hue_api

type GroupAction struct {
	On               bool      `json:"on"`
	Brightness       int       `json:"bri"`
	Hue              int       `json:"hue"`
	Saturation       int       `json:"sat"`
	Effect           string    `json:"effect"`
	XY               []float32 `json:"xy"`
	ColorTemperature int       `json:"ct"`
	Alert            string    `json:"alert"`
	ColorMode        string    `json:"colormode"`
}
