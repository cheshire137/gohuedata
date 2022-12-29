package data_store

import "github.com/cheshire137/gohuedata/pkg/hue_api"

func (ds *DataStore) AddLight(bridge *hue_api.Bridge, id string, light *hue_api.Light) error {
	insertQuery := `INSERT INTO lights (unique_id, id, name, bridge_ip_address) VALUES (?, ?, ?, ?)
		ON CONFLICT(unique_id)
		DO UPDATE SET id = excluded.id, name = excluded.name, bridge_ip_address = excluded.bridge_ip_address`
	stmt, err := ds.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(light.UniqueID, id, light.Name, bridge.IPAddress)
	if err != nil {
		return err
	}
	return nil
}
