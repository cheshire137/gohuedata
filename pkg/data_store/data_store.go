package data_store

import "database/sql"

type DataStore struct {
	db *sql.DB
}

func NewDataStore(db *sql.DB) (*DataStore, error) {
	err := createTables(db)
	if err != nil {
		return nil, err
	}
	return &DataStore{db: db}, nil
}

func createTables(db *sql.DB) error {
	err := createTemperatureSensorsTable(db)
	if err != nil {
		return err
	}
	err = createTemperatureReadingsTable(db)
	if err != nil {
		return err
	}
	return createHueBridgesTable(db)
}
