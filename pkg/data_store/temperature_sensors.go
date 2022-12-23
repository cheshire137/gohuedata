package data_store

import (
	"database/sql"

	"github.com/cheshire137/gohuedata/pkg/hueapi"
)

func (ds *DataStore) addTemperatureSensor(sensor *hueapi.TemperatureSensor) error {
	insertQuery := `INSERT INTO temperature_sensors (id, name) VALUES (?, ?)
		ON CONFLICT(id) DO UPDATE SET name = excluded.name`
	stmt, err := ds.db.Prepare(insertQuery)
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
