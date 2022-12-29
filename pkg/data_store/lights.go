package data_store

import "github.com/cheshire137/gohuedata/pkg/hue_api"

func (ds *DataStore) addLight(bridge *hue_api.Bridge, light *hue_api.Light) error {
	insertQuery := `INSERT INTO lights (id, name, bridge_ip_address) VALUES (?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET name = excluded.name`
	stmt, err := ds.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(light.UniqueID, light.Name, bridge.IPAddress)
	if err != nil {
		return err
	}
	return nil
}
