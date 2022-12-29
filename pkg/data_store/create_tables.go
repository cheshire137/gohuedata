package data_store

func (ds *DataStore) CreateTables() error {
	err := ds.createTemperatureSensorsTable()
	if err != nil {
		return err
	}
	err = ds.createTemperatureReadingsTable()
	if err != nil {
		return err
	}
	err = ds.createLightsTable()
	if err != nil {
		return err
	}
	return ds.createHueBridgesTable()
}

func (ds *DataStore) createTemperatureReadingsTable() error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS temperature_readings (
		temperature_sensor_id TEXT NOT NULL,
		last_updated TEXT NOT NULL,
		temperature REAL NOT NULL,
		units TEXT NOT NULL,
		PRIMARY KEY (temperature_sensor_id, last_updated, temperature, units)
	)`
	stmt, err := ds.db.Prepare(createTableQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStore) createHueBridgesTable() error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS hue_bridges (
		ip_address TEXT PRIMARY KEY,
		name TEXT NOT NULL
	)`
	stmt, err := ds.db.Prepare(createTableQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStore) createLightsTable() error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS lights (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		bridge_ip_address TEXT NOT NULL
	)`
	stmt, err := ds.db.Prepare(createTableQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	createIndexQuery := `CREATE UNIQUE INDEX IF NOT EXISTS idx_lights_id_bridge_ip ON lights (id, bridge_ip_address)`
	stmt, err = ds.db.Prepare(createIndexQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStore) createTemperatureSensorsTable() error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS temperature_sensors (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		bridge_ip_address TEXT NOT NULL
	)`
	stmt, err := ds.db.Prepare(createTableQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	createIndexQuery := `CREATE UNIQUE INDEX IF NOT EXISTS idx_temperature_sensors_id_bridge_ip
		ON temperature_sensors (id, bridge_ip_address)`
	stmt, err = ds.db.Prepare(createIndexQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
