package data_store

import (
	"time"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

func (ds *DataStore) AddLightState(bridge *hue_api.Bridge, light *hue_api.Light) error {
	insertQuery := `INSERT INTO light_states (light_unique_id, timestamp, on)
		VALUES (?, ?, ?) ON CONFLICT(light_unique_id, timestamp) DO UPDATE SET on = excluded.on`
	stmt, err := ds.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	timestamp := time.Now().Format(time.RFC3339)
	_, err = stmt.Exec(light.UniqueID, timestamp, light.State.On)
	if err != nil {
		return err
	}
	return nil
}
