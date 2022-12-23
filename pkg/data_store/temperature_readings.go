package data_store

import (
	"database/sql"

	"github.com/cheshire137/gohuedata/pkg/hueapi"
)

func (ds *DataStore) AddTemperatureReading(sensor *hueapi.TemperatureSensor) error {
	err := ds.addTemperatureSensor(sensor)
	if err != nil {
		return err
	}
	insertQuery := `INSERT INTO temperature_readings (temperature_sensor_id, last_updated, fahrenheit_temperature)
		VALUES (?, ?, ?)`
	stmt, err := ds.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sensor.UniqueID, sensor.State.LastUpdated, sensor.State.FahrenheitTemperature())
	if err != nil {
		return err
	}
	return nil
}

func createTemperatureReadingsTable(db *sql.DB) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS temperature_readings (
		temperature_sensor_id TEXT NOT NULL,
		last_updated TEXT NOT NULL,
		fahrenheit_temperature REAL NOT NULL,
		PRIMARY KEY (temperature_sensor_id, last_updated, fahrenheit_temperature)
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
