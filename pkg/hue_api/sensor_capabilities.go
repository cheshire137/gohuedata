package hue_api

type SensorCapabilities struct {
	Certified bool          `json:"certified"`
	Primary   bool          `json:"primary"`
	Inputs    []SensorInput `json:"inputs"`
}
