package temperature_sensor_store

import "database/sql"

type TemperatureSensorStore struct {
	db *sql.DB
}

func NewTemperatureSensorStore(db *sql.DB) (*TemperatureSensorStore, error) {
	createTableQuery := `CREATE TABLE IF NOT EXISTS temperature_sensors (
		id TEXT PRIMARY KEY,
		name TEXT,
		last_updated TEXT,
		fahrenheit_temperature REAL
	)`
	stmt, err := db.Prepare(createTableQuery)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}
	return &TemperatureSensorStore{db: db}, nil
}
