package sensor_loader

import (
	"fmt"
	"time"

	options "github.com/cheshire137/gohuedata/pkg/cli_options"
	"github.com/cheshire137/gohuedata/pkg/data_store"
	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type SensorLoader struct {
	SensorSelection    options.SensorSelection
	MotionSensors      []*hue_api.MotionSensor
	TemperatureSensors []*hue_api.TemperatureSensor
	OtherSensors       []*hue_api.Sensor
}

func NewSensorLoader(hueClient *hue_api.Client, sensorSelection options.SensorSelection) (*SensorLoader, error) {
	sensors, err := hueClient.GetSensors(sensorSelection)
	if err != nil {
		return nil, err
	}

	tempSensors := []*hue_api.TemperatureSensor{}
	motionSensors := []*hue_api.MotionSensor{}
	otherSensors := []*hue_api.Sensor{}

	for _, sensor := range sensors {
		if sensorSelection != options.Motion {
			tempSensor, ok := sensor.(*hue_api.TemperatureSensor)
			if ok {
				tempSensors = append(tempSensors, tempSensor)
				continue
			}
		}

		if sensorSelection != options.Temperature {
			motionSensor, ok := sensor.(*hue_api.MotionSensor)
			if ok {
				motionSensors = append(motionSensors, motionSensor)
				continue
			}
		}

		if sensorSelection == options.AllSensors {
			sensor, ok := sensor.(*hue_api.Sensor)
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

func (sl *SensorLoader) DisplaySensors(quietMode bool) {
	displayAllSensors := sl.SensorSelection == options.AllSensors

	if displayAllSensors {
		totalSensors := sl.TotalSensors()
		units := util.Pluralize(totalSensors, "sensor", "sensors")
		util.LogSuccess("Got %d %s%s", totalSensors, units, util.LinePunctuation(quietMode))

		if !quietMode {
			for i, sensor := range sl.OtherSensors {
				fmt.Printf("%d. %s\n", i+1, sensor.String())
			}
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
		util.LogSuccess("%s %d temperature %s%s", intro, totalTempSensors, units, util.LinePunctuation(quietMode))

		if !quietMode {
			for i, sensor := range sl.TemperatureSensors {
				fmt.Printf("%d. %s\n", i+1, sensor.String())
			}
		}
	}

	totalMotionSensors := sl.TotalMotionSensors()
	if totalMotionSensors > 0 {
		units := util.Pluralize(totalMotionSensors, "sensor", "sensors")
		util.LogSuccess("%s %d motion %s%s", intro, totalMotionSensors, units, util.LinePunctuation(quietMode))

		if !quietMode {
			for i, sensor := range sl.MotionSensors {
				fmt.Printf("%d. %s\n", i+1, sensor.String())
			}
		}
	}
}

func (sl *SensorLoader) AllSensors() []*hue_api.Sensor {
	allSensors := make([]*hue_api.Sensor, sl.TotalSensors())
	start := 0

	for i, sensor := range sl.OtherSensors {
		allSensors[i+start] = sensor
	}
	start += len(sl.OtherSensors)

	for i, tempSensor := range sl.TemperatureSensors {
		allSensors[i+start] = tempSensor.ToSensor()
	}
	start += len(sl.TemperatureSensors)

	for i, motionSensor := range sl.MotionSensors {
		allSensors[i+start] = motionSensor.ToSensor()
	}

	return allSensors
}

func (sl *SensorLoader) SensorStates() []*hue_api.SensorState {
	sensorStates := make([]*hue_api.SensorState, sl.TotalSensors())
	for i, sensor := range sl.AllSensors() {
		sensorStates[i] = &sensor.State
	}
	return sensorStates
}

func (sl *SensorLoader) MostRecentlyUpdatedSensorState() *hue_api.SensorState {
	var latestLastUpdated *time.Time
	var latestSensorState *hue_api.SensorState
	for _, sensorState := range sl.SensorStates() {
		lastUpdatedAt, err := sensorState.LastUpdatedAt()
		if err != nil {
			continue
		}
		if latestLastUpdated == nil || lastUpdatedAt.After(*latestLastUpdated) {
			latestLastUpdated = lastUpdatedAt
			latestSensorState = sensorState
		}
	}
	return latestSensorState
}

func (sl *SensorLoader) SaveTemperatureSensorReadings(bridge *hue_api.Bridge, dataStore *data_store.DataStore, fahrenheit bool) error {
	for _, sensor := range sl.TemperatureSensors {
		err := dataStore.AddTemperatureReading(bridge, sensor, fahrenheit)
		if err != nil {
			return err
		}
	}
	return nil
}
