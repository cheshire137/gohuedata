package hue_api

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
	Fahrenheit       bool
}

func NewTemperatureSensor(s Sensor, fahrenheit bool) *TemperatureSensor {
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
		Fahrenheit:       fahrenheit,
	}
}

func (s *TemperatureSensor) Temperature() float32 {
	if s.Fahrenheit {
		return s.State.FahrenheitTemperature()
	}
	return s.State.CelsiusTemperature()
}

func (s *TemperatureSensor) TempUnits() string {
	if s.Fahrenheit {
		return "°F"
	}
	return "°C"
}

func (s *TemperatureSensor) String() string {
	lastUpdatedSummary := s.State.LastUpdatedSummary()
	if lastUpdatedSummary == "" {
		return fmt.Sprintf("%s -- %.1f%s", s.Name, s.Temperature(), s.TempUnits())
	}
	return fmt.Sprintf("%s -- %.1f%s as of %s", s.Name, s.Temperature(), s.TempUnits(), lastUpdatedSummary)
}

func (s *TemperatureSensor) ToSensor() *Sensor {
	return &Sensor{
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
