package data_store

import (
	"database/sql"

	"github.com/cheshire137/gohuedata/pkg/hueapi"
)

type TemperatureSensorStore struct {
	db *sql.DB
}

func NewTemperatureSensorStore(db *sql.DB) (*TemperatureSensorStore, error) {
	err := createTemperatureSensorsTable(db)
	if err != nil {
		return nil, err
	}
	err = createTemperatureReadingsTable(db)
	if err != nil {
		return nil, err
	}
	return &TemperatureSensorStore{db: db}, nil
}

func (tss *TemperatureSensorStore) AddTemperatureReading(sensor *hueapi.TemperatureSensor) error {
	err := tss.addTemperatureSensor(sensor)
	if err != nil {
		return err
	}
	insertQuery := `INSERT INTO temperature_readings (temperature_sensor_id, last_updated, fahrenheit_temperature)
		VALUES (?, ?, ?)`
	stmt, err := tss.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sensor.UniqueID, sensor.State.LastUpdated, sensor.State.FahrenheitTemperature())
	if err != nil {
		return err
	}
	return nil
}

func (tss *TemperatureSensorStore) addTemperatureSensor(sensor *hueapi.TemperatureSensor) error {
	insertQuery := `INSERT INTO temperature_sensors (id, name) VALUES (?, ?)
		ON CONFLICT(id) DO UPDATE SET name = excluded.name`
	stmt, err := tss.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sensor.UniqueID, sensor.Name)
	if err != nil {
		return err
	}
	return nil
}

func createTemperatureSensorsTable(db *sql.DB) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS temperature_sensors (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL
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
