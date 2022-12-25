package data_store

import (
	"strings"

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

func (ds *DataStore) LoadHueBridgesByIPAddress(ipAddresses ...string) ([]*HueBridge, error) {
	bridges := []*HueBridge{}
	if len(ipAddresses) == 0 {
		return bridges, nil
	}
	queryStr := `SELECT ip_address, name FROM hue_bridges WHERE ip_address IN (?` +
		strings.Repeat(", ?", len(ipAddresses)-1) + `)`
	rows, err := ds.db.Query(queryStr, ipAddresses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		bridge := &HueBridge{}
		err = rows.Scan(&bridge.IPAddress, &bridge.Name)
		if err != nil {
			return nil, err
		}
		bridges = append(bridges, bridge)
	}
	return bridges, nil
}
