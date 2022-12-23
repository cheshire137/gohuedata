package data_store

import "database/sql"

type DataStore struct {
	db *sql.DB
}

func NewDataStore(db *sql.DB) *DataStore {
	return &DataStore{db: db}
}

func (ds *DataStore) CreateTables() error {
	err := createTemperatureSensorsTable(ds.db)
	if err != nil {
		return err
	}
	err = createTemperatureReadingsTable(ds.db)
	if err != nil {
		return err
	}
	return createHueBridgesTable(ds.db)
}
