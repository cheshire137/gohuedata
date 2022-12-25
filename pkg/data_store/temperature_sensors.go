package data_store

import (
	"github.com/cheshire137/gohuedata/pkg/hue_api"
)

func (ds *DataStore) addTemperatureSensor(bridge *hue_api.Bridge, sensor *hue_api.TemperatureSensor) error {
	insertQuery := `INSERT INTO temperature_sensors (id, name, bridge_ip_address) VALUES (?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET name = excluded.name`
	stmt, err := ds.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sensor.UniqueID, sensor.Name, bridge.IPAddress)
	if err != nil {
		return err
	}
	return nil
}
