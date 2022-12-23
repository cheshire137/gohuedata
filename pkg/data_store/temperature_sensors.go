package data_store

import (
	"database/sql"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

func (ds *DataStore) addTemperatureSensor(bridge *hue_api.Bridge, sensor *hue_api.TemperatureSensor) error {
	insertQuery := `INSERT INTO temperature_sensors (id, name, bridge_ip_address) VALUES (?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET name = excluded.name`
	stmt, err := ds.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sensor.UniqueID, sensor.Name, bridge.IPAddress)
	if err != nil {
		return err
	}
	return nil
}

func createTemperatureSensorsTable(db *sql.DB) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS temperature_sensors (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		bridge_ip_address TEXT NOT NULL
	)`
	stmt, err := db.Prepare(createTableQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	createIndexQuery := `CREATE UNIQUE INDEX IF NOT EXISTS idx_temperature_sensors_id_bridge_ip
		ON temperature_sensors (id, bridge_ip_address)`
	stmt, err = db.Prepare(createIndexQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
