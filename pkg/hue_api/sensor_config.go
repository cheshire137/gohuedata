package hue_api

type SensorConfig struct {
	On              bool   `json:"on"`
	Longitude       string `json:"long"`
	Latitude        string `json:"lat"`
	SunriseOffset   int    `json:"sunriseoffset"`
	SunsetOffset    int    `json:"sunsetoffset"`
	Reachable       bool   `json:"reachable"`
	UserTest        bool   `json:"usertest"`
	LEDIndication   bool   `json:"ledindication"`
	Battery         int    `json:"battery"`
	Alert           string `json:"alert"`
	ThresholdDark   int    `json:"tholddark"`
	ThresholdOffset int    `json:"tholdoffset"`
	Sensitivity     int    `json:"sensitivity"`
	SensitivityMax  int    `json:"sensitivitymax"`
}
