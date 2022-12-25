package data_store

import (
	"strings"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/pagination"
)

type TemperatureSensorFilter struct {
	Page       int
	PerPage    int
	BridgeName string
}

const temperatureSensorJoins = `
	INNER JOIN hue_bridges ON temperature_sensors.bridge_ip_address = hue_bridges.ip_address `

func (ds *DataStore) TotalTemperatureSensors(filter *TemperatureSensorFilter) (int, error) {
	whereClause, values := buildTemperatureSensorWhereConditions(filter)
	queryStr := "SELECT COUNT(*) FROM temperature_sensors" + temperatureSensorJoins + whereClause
	var count int
	err := ds.db.QueryRow(queryStr, values...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (ds *DataStore) LoadTemperatureSensors(filter *TemperatureSensorFilter) ([]*TemperatureSensor, error) {
	limit, offset := temperatureSensorLimitAndOffset(filter)
	whereClause, values := buildTemperatureSensorWhereConditions(filter)
	queryStr := `SELECT temperature_sensors.id AS sensor_id,
			temperature_sensors.name AS sensor_name,
			temperature_sensors.bridge_ip_address AS bridge_ip_address,
			hue_bridges.name AS bridge_name,
			MAX(temperature_readings.last_updated) AS sensor_last_updated
		FROM temperature_sensors` + temperatureSensorJoins + `
		LEFT OUTER JOIN temperature_readings ON temperature_sensors.id = temperature_readings.temperature_sensor_id
		` + whereClause + `
		GROUP BY temperature_sensors.id, temperature_sensors.name, temperature_sensors.bridge_ip_address, hue_bridges.name
		ORDER BY temperature_sensors.name ASC, hue_bridges.name ASC
		LIMIT ? OFFSET ?`
	values = append(values, limit, offset)

	rows, err := ds.db.Query(queryStr, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*TemperatureSensor
	bridgesByIPAddress := map[string]*HueBridge{}

	for rows.Next() {
		var sensor TemperatureSensor
		var bridgeName string

		err = rows.Scan(&sensor.ID, &sensor.Name, &sensor.bridgeIPAddress, &bridgeName, &sensor.LastUpdated)
		if err != nil {
			return nil, err
		}

		bridge, ok := bridgesByIPAddress[sensor.bridgeIPAddress]
		if !ok {
			bridge = &HueBridge{
				ID:        sensor.bridgeIPAddress,
				IPAddress: sensor.bridgeIPAddress,
				Name:      bridgeName,
			}
			bridgesByIPAddress[sensor.bridgeIPAddress] = bridge
		}

		sensor.Bridge = bridge
		sensors = append(sensors, &sensor)
	}

	return sensors, nil
}

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

func temperatureSensorLimitAndOffset(filter *TemperatureSensorFilter) (int, int) {
	page, perPage := 1, 10
	if filter != nil {
		page, perPage = pagination.ConstrainPageAndPerPage(filter.Page, filter.PerPage)
	}
	return pagination.GetLimitAndOffset(page, perPage)
}

func buildTemperatureSensorWhereConditions(filter *TemperatureSensorFilter) (string, []interface{}) {
	if filter == nil {
		return "", []interface{}{}
	}

	conditions := []string{}
	values := []interface{}{}

	if filter.BridgeName != "" {
		conditions = append(conditions, "LOWER(hue_bridges.name) = ?")
		values = append(values, strings.ToLower(filter.BridgeName))
	}

	if len(conditions) == 0 {
		return "", []interface{}{}
	}

	whereClause := "WHERE " + strings.Join(conditions, " AND ")
	return whereClause, values
}

func setLastUpdatedOnSensors(sensors []*TemperatureSensor, readings []*TemperatureReading) {
	for _, sensor := range sensors {
		for _, reading := range readings {
			if reading.temperatureSensorID == sensor.ID && (sensor.LastUpdated == "" || reading.Timestamp > sensor.LastUpdated) {
				sensor.LastUpdated = reading.Timestamp
			}
		}
	}
}
