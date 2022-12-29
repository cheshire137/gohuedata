package data_store

import "database/sql"

func (ds *DataStore) CreateTables() error {
	err := createTemperatureSensorsTable(ds.db)
	if err != nil {
		return err
	}
	err = createTemperatureReadingsTable(ds.db)
	if err != nil {
		return err
	}
	err = createLightsTable(ds.db)
	if err != nil {
		return err
	}
	return createHueBridgesTable(ds.db)
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

func createHueBridgesTable(db *sql.DB) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS hue_bridges (
		ip_address TEXT PRIMARY KEY,
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

func createLightsTable(db *sql.DB) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS lights (
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
	createIndexQuery := `CREATE UNIQUE INDEX IF NOT EXISTS idx_lights_id_bridge_ip ON lights (id, bridge_ip_address)`
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
