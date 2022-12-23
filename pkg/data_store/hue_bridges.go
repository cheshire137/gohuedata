package data_store

import (
	"database/sql"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

func (ds *DataStore) AddHueBridge(bridge *hue_api.Bridge) error {
	insertQuery := `INSERT INTO hue_bridges (ip_address, name) VALUES (?, ?)
		ON CONFLICT(ip_address) DO UPDATE SET name = excluded.name`
	stmt, err := ds.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(bridge.IPAddress, bridge.Name)
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
