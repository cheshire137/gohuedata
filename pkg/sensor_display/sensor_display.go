package sensor_display

import (
	"fmt"

	"github.com/cheshire137/gohuedata/pkg/hueapi"
	"github.com/cheshire137/gohuedata/pkg/options"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type SensorDisplay struct {
	SensorSelection    options.SensorSelection
	MotionSensors      []*hueapi.MotionSensor
	TemperatureSensors []*hueapi.TemperatureSensor
	OtherSensors       []*hueapi.Sensor
}

func NewSensorDisplay(hueClient *hueapi.Client, sensorSelection options.SensorSelection) (*SensorDisplay, error) {
	sensors, err := hueClient.GetSensors(sensorSelection)
	if err != nil {
		return nil, err
	}

	tempSensors := []*hueapi.TemperatureSensor{}
	motionSensors := []*hueapi.MotionSensor{}
	otherSensors := []*hueapi.Sensor{}

	for _, sensor := range sensors {
		if sensorSelection != options.Motion {
			tempSensor, ok := sensor.(*hueapi.TemperatureSensor)
			if ok {
				tempSensors = append(tempSensors, tempSensor)
				continue
			}
		}

		if sensorSelection != options.Temperature {
			motionSensor, ok := sensor.(*hueapi.MotionSensor)
			if ok {
				motionSensors = append(motionSensors, motionSensor)
				continue
			}
		}

		if sensorSelection == options.AllSensors {
			sensor, ok := sensor.(*hueapi.Sensor)
			if !ok {
				return nil, fmt.Errorf("Unknown sensor type: %v", sensor)
			}
			otherSensors = append(otherSensors, sensor)
		}
	}

	return &SensorDisplay{
		SensorSelection:    sensorSelection,
		MotionSensors:      motionSensors,
		TemperatureSensors: tempSensors,
		OtherSensors:       otherSensors,
	}, nil
}

func (sd *SensorDisplay) TotalSensors() int {
	return sd.TotalMotionSensors() + sd.TotalTemperatureSensors() + len(sd.OtherSensors)
}

func (sd *SensorDisplay) TotalTemperatureSensors() int {
	return len(sd.TemperatureSensors)
}

func (sd *SensorDisplay) TotalMotionSensors() int {
	return len(sd.MotionSensors)
}

func (sd *SensorDisplay) DisplaySensors() {
	displayAllSensors := sd.SensorSelection == options.AllSensors

	if displayAllSensors {
		totalSensors := sd.TotalSensors()
		units := util.Pluralize(totalSensors, "sensor", "sensors")
		fmt.Printf("\n✅ Got %d %s:\n", totalSensors, units)

		for i, sensor := range sd.OtherSensors {
			fmt.Printf("%d. %s\n", i+1, sensor.String())
		}
	}

	var intro string
	if displayAllSensors {
		intro = "Including"
	} else {
		intro = "Got"
	}

	totalTempSensors := sd.TotalTemperatureSensors()
	if totalTempSensors > 0 {
		units := util.Pluralize(totalTempSensors, "sensor", "sensors")
		fmt.Printf("\n✅ %s %d temperature %s:\n", intro, totalTempSensors, units)
		for i, sensor := range sd.TemperatureSensors {
			fmt.Printf("%d. %s\n", i+1, sensor.String())
		}
	}

	totalMotionSensors := sd.TotalMotionSensors()
	if totalMotionSensors > 0 {
		units := util.Pluralize(totalMotionSensors, "sensor", "sensors")
		fmt.Printf("\n✅ %s %d motion %s:\n", intro, totalMotionSensors, units)
		for i, sensor := range sd.MotionSensors {
			fmt.Printf("%d. %s\n", i+1, sensor.String())
		}
	}
}
