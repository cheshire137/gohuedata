package data_store

import (
	"database/sql"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

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
