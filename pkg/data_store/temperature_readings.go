package data_store

import (
	"database/sql"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

type TemperatureReading struct {
	TemperatureSensorID string  `json:"temperatureSensorID"`
	LastUpdated         string  `json:"lastUpdated"`
	Temperature         float32 `json:"temperature"`
	Units               string  `json:"units"`
}

func (ds *DataStore) LoadTemperatureReadings(page int) ([]*TemperatureReading, error) {
	perPage := 100
	rows, err := ds.db.Query(`SELECT temperature_sensor_id, last_updated, temperature, units
		FROM temperature_readings
		ORDER BY last_updated DESC
		LIMIT ? OFFSET ?`, perPage, (page-1)*perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var readings []*TemperatureReading
	for rows.Next() {
		var reading TemperatureReading
		err = rows.Scan(&reading.TemperatureSensorID, &reading.LastUpdated, &reading.Temperature, &reading.Units)
		if err != nil {
			return nil, err
		}
		readings = append(readings, &reading)
	}
	return readings, nil
}

func (ds *DataStore) AddTemperatureReading(bridge *hue_api.Bridge, sensor *hue_api.TemperatureSensor, fahrenheit bool) error {
	err := ds.addTemperatureSensor(bridge, sensor)
	if err != nil {
		return err
	}
	insertQuery := `INSERT INTO temperature_readings (temperature_sensor_id, last_updated, temperature, units)
		VALUES (?, ?, ?, ?) ON CONFLICT DO NOTHING`
	stmt, err := ds.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	var temperature float32
	var units string
	if fahrenheit {
		temperature = sensor.State.FahrenheitTemperature()
		units = "F"
	} else {
		temperature = sensor.State.CelsiusTemperature()
		units = "C"
	}
	_, err = stmt.Exec(sensor.UniqueID, sensor.State.LastUpdated, temperature, units)
	if err != nil {
		return err
	}
	return nil
}

func createTemperatureReadingsTable(db *sql.DB) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS temperature_readings (
		temperature_sensor_id TEXT NOT NULL,
		last_updated TEXT NOT NULL,
		temperature REAL NOT NULL,
		units TEXT NOT NULL,
		PRIMARY KEY (temperature_sensor_id, last_updated, temperature, units)
	)`
	stmt, err := db.Prepare(createTableQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
