package hueapi

import "fmt"

// https://developers.meethue.com/develop/hue-api/5-sensors-api/
type Sensor struct {
	State            SensorState        `json:"state"`
	Config           SensorConfig       `json:"config"`
	Name             string             `json:"name"`
	Type             string             `json:"type"`
	ModelID          string             `json:"modelid"`
	ManufacturerName string             `json:"manufacturername"`
	SoftwareVersion  string             `json:"swversion"`
	UniqueID         string             `json:"uniqueid"`
	DiversityID      string             `json:"diversityid"`
	ProductName      string             `json:"productname"`
	Capabilities     SensorCapabilities `json:"capabilities"`
	Recycle          bool               `json:"recycle"`
}

func (s *Sensor) String() string {
	return fmt.Sprintf("%s -- %s", s.Name, s.Type)
}
