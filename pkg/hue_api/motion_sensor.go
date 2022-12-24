package hue_api

import "fmt"

type MotionSensor struct {
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

func NewMotionSensor(s Sensor) *MotionSensor {
	return &MotionSensor{
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

func (s *MotionSensor) String() string {
	lastUpdatedSummary := s.State.LastUpdatedSummary()
	if lastUpdatedSummary != "" {
		return fmt.Sprintf("%s -- %s", s.Name, lastUpdatedSummary)
	}
	return s.Name
}

func (s *MotionSensor) ToSensor() *Sensor {
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
