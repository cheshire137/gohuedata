package hueapi

import "fmt"

type TemperatureSensor struct {
	State            SensorState
	Config           SensorConfig
	Name             string
	ModelID          string
	ManufacturerName string
	SoftwareVersion  string
	UniqueID         string
	DiversityID      string
	ProductName      string
	Capabilities     SensorCapabilities
	Recycle          bool
}

func NewTemperatureSensor(s Sensor) *TemperatureSensor {
	return &TemperatureSensor{
		State:            s.State,
		Config:           s.Config,
		Name:             s.Name,
		ModelID:          s.ModelID,
		ManufacturerName: s.ManufacturerName,
		SoftwareVersion:  s.SoftwareVersion,
		UniqueID:         s.UniqueID,
		DiversityID:      s.DiversityID,
		ProductName:      s.ProductName,
		Capabilities:     s.Capabilities,
		Recycle:          s.Recycle,
	}
}

func (s *TemperatureSensor) String() string {
	lastUpdatedSummary := s.State.LastUpdatedSummary()
	if lastUpdatedSummary == "" {
		return fmt.Sprintf("%s -- %d°F", s.Name, s.State.FahrenheitTemperature())
	}
	return fmt.Sprintf("%s -- %d°F as of %s", s.Name, s.State.FahrenheitTemperature(), lastUpdatedSummary)
}
