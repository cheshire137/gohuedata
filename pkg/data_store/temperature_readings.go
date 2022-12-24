package data_store

import (
	"database/sql"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

type TemperatureReading struct {
	TemperatureSensor *TemperatureSensor `json:"temperatureSensor"`
	LastUpdated       string             `json:"lastUpdated"`
	Temperature       float32            `json:"temperature"`
	Units             string             `json:"units"`
}

type TemperatureReadingFilter struct {
	Page    int
	PerPage int
}

func (ds *DataStore) TotalTemperatureReadings(filter *TemperatureReadingFilter) (int, error) {
	var count int
	query := `SELECT COUNT(*)
		FROM temperature_readings
		INNER JOIN temperature_sensors ON temperature_readings.temperature_sensor_id = temperature_sensors.id
		INNER JOIN hue_bridges ON temperature_sensors.bridge_ip_address = hue_bridges.ip_address`
	err := ds.db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (ds *DataStore) LoadTemperatureReadings(filter *TemperatureReadingFilter) ([]*TemperatureReading, error) {
	var page int
	if filter == nil {
		page = 1
	} else {
		page = filter.Page
	}
	var perPage int
	if filter == nil {
		perPage = 10
	} else {
		perPage = filter.PerPage
	}
	rows, err := ds.db.Query(`SELECT temperature_readings.last_updated,
			temperature_readings.temperature,
			temperature_readings.units,
			temperature_sensors.name AS sensor_name,
			hue_bridges.name AS bridge_name
		FROM temperature_readings
		INNER JOIN temperature_sensors ON temperature_readings.temperature_sensor_id = temperature_sensors.id
		INNER JOIN hue_bridges ON temperature_sensors.bridge_ip_address = hue_bridges.ip_address
		ORDER BY temperature_readings.last_updated DESC, temperature_sensors.name ASC, hue_bridges.name ASC
		LIMIT ? OFFSET ?`, perPage, (page-1)*perPage)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var readings []*TemperatureReading
	for rows.Next() {
		var reading TemperatureReading
		var sensor TemperatureSensor
		var bridge HueBridge
		err = rows.Scan(&reading.LastUpdated, &reading.Temperature, &reading.Units, &sensor.Name, &bridge.Name)
		if err != nil {
			return nil, err
		}
		sensor.Bridge = &bridge
		reading.TemperatureSensor = &sensor
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
