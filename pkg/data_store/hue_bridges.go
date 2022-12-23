package data_store

import (
	"database/sql"
)

func createHueBridgesTable(db *sql.DB) error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS hue_bridges (
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
