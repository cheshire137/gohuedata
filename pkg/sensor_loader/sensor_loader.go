package sensor_loader

import (
	"fmt"

	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/hueapi"
	"github.com/cheshire137/gohuedata/pkg/options"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type SensorLoader struct {
	SensorSelection    options.SensorSelection
	MotionSensors      []*hueapi.MotionSensor
	TemperatureSensors []*hueapi.TemperatureSensor
	OtherSensors       []*hueapi.Sensor
}

func NewSensorLoader(hueClient *hueapi.Client, sensorSelection options.SensorSelection) (*SensorLoader, error) {
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

	return &SensorLoader{
		SensorSelection:    sensorSelection,
		MotionSensors:      motionSensors,
		TemperatureSensors: tempSensors,
		OtherSensors:       otherSensors,
	}, nil
}

func (sl *SensorLoader) TotalSensors() int {
	return sl.TotalMotionSensors() + sl.TotalTemperatureSensors() + len(sl.OtherSensors)
}

func (sl *SensorLoader) TotalTemperatureSensors() int {
	return len(sl.TemperatureSensors)
}

func (sl *SensorLoader) TotalMotionSensors() int {
	return len(sl.MotionSensors)
}

func (sl *SensorLoader) DisplaySensors() {
	displayAllSensors := sl.SensorSelection == options.AllSensors

	if displayAllSensors {
		totalSensors := sl.TotalSensors()
		units := util.Pluralize(totalSensors, "sensor", "sensors")
		util.LogSuccess("Got %d %s:", totalSensors, units)

		for i, sensor := range sl.OtherSensors {
			fmt.Printf("%d. %s\n", i+1, sensor.String())
		}
	}

	var intro string
	if displayAllSensors {
		intro = "Including"
	} else {
		intro = "Got"
	}

	totalTempSensors := sl.TotalTemperatureSensors()
	if totalTempSensors > 0 {
		units := util.Pluralize(totalTempSensors, "sensor", "sensors")
		util.LogSuccess("%s %d temperature %s:", intro, totalTempSensors, units)
		for i, sensor := range sl.TemperatureSensors {
			fmt.Printf("%d. %s\n", i+1, sensor.String())
		}
	}

	totalMotionSensors := sl.TotalMotionSensors()
	if totalMotionSensors > 0 {
		units := util.Pluralize(totalMotionSensors, "sensor", "sensors")
		util.LogSuccess("%s %d motion %s:", intro, totalMotionSensors, units)
		for i, sensor := range sl.MotionSensors {
			fmt.Printf("%d. %s\n", i+1, sensor.String())
		}
	}
}

func (sl *SensorLoader) SaveTemperatureSensorReadings(tss *data_store.DataStore) error {
	for _, sensor := range sl.TemperatureSensors {
		err := tss.AddTemperatureReading(sensor)
		if err != nil {
			return err
		}
	}
	return nil
}
