package data_store

import (
	"strings"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/pagination"
)

type TemperatureReading struct {
	TemperatureSensor *TemperatureSensor `json:"temperatureSensor"`
	LastUpdated       string             `json:"lastUpdated"`
	Temperature       float32            `json:"temperature"`
	Units             string             `json:"units"`
}

type TemperatureReadingFilter struct {
	Page          int
	PerPage       int
	BridgeName    string
	UpdatedSince  string
	UpdatedBefore string
}

const temperatureReadingJoins = `
	INNER JOIN temperature_sensors ON temperature_readings.temperature_sensor_id = temperature_sensors.id
	INNER JOIN hue_bridges ON temperature_sensors.bridge_ip_address = hue_bridges.ip_address `

func (ds *DataStore) TotalTemperatureReadings(filter *TemperatureReadingFilter) (int, error) {
	whereClause, values := buildTemperatureReadingWhereConditions(filter)
	queryStr := "SELECT COUNT(*) FROM temperature_readings" + temperatureReadingJoins + whereClause
	var count int
	err := ds.db.QueryRow(queryStr, values...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (ds *DataStore) LoadTemperatureReadings(filter *TemperatureReadingFilter) ([]*TemperatureReading, error) {
	limit, offset := temperatureReadingLimitAndOffset(filter)
	whereClause, values := buildTemperatureReadingWhereConditions(filter)
	queryStr := `SELECT temperature_readings.last_updated,
			temperature_readings.temperature,
			temperature_readings.units,
			temperature_sensors.name AS sensor_name,
			hue_bridges.name AS bridge_name
		FROM temperature_readings` + temperatureReadingJoins + whereClause + `
		ORDER BY temperature_readings.last_updated DESC, temperature_sensors.name ASC, hue_bridges.name ASC
		LIMIT ? OFFSET ?`
	values = append(values, limit, offset)

	rows, err := ds.db.Query(queryStr, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readings []*TemperatureReading
	for rows.Next() {
		var reading TemperatureReading
		var sensor TemperatureSensor
		var bridge HueBridge

		err = rows.Scan(&reading.LastUpdated, &reading.Temperature, &reading.Units, &sensor.Name, &bridge.Name)
		if err != nil {
			return nil, err
		}

		sensor.Bridge = &bridge
		reading.TemperatureSensor = &sensor
		readings = append(readings, &reading)
	}

	return readings, nil
}

func (ds *DataStore) AddTemperatureReading(bridge *hue_api.Bridge, sensor *hue_api.TemperatureSensor, fahrenheit bool) error {
	err := ds.addTemperatureSensor(bridge, sensor)
	if err != nil {
		return err
	}
	insertQuery := `INSERT INTO temperature_readings (temperature_sensor_id, last_updated, temperature, units)
		VALUES (?, ?, ?, ?) ON CONFLICT DO NOTHING`
	stmt, err := ds.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	var temperature float32
	var units string
	if fahrenheit {
		temperature = sensor.State.FahrenheitTemperature()
		units = "F"
	} else {
		temperature = sensor.State.CelsiusTemperature()
		units = "C"
	}
	_, err = stmt.Exec(sensor.UniqueID, sensor.State.LastUpdated, temperature, units)
	if err != nil {
		return err
	}
	return nil
}

func buildTemperatureReadingWhereConditions(filter *TemperatureReadingFilter) (string, []interface{}) {
	if filter == nil {
		return "", []interface{}{}
	}

	conditions := []string{}
	values := []interface{}{}

	if filter.BridgeName != "" {
		conditions = append(conditions, "LOWER(hue_bridges.name) = ?")
		values = append(values, strings.ToLower(filter.BridgeName))
	}

	if filter.UpdatedSince != "" {
		conditions = append(conditions, "temperature_readings.last_updated >= ?")
		values = append(values, filter.UpdatedSince)
	}

	if filter.UpdatedBefore != "" {
		conditions = append(conditions, "temperature_readings.last_updated < ?")
		values = append(values, filter.UpdatedBefore)
	}

	if len(conditions) == 0 {
		return "", []interface{}{}
	}

	whereClause := "WHERE " + strings.Join(conditions, " AND ")
	return whereClause, values
}

func temperatureReadingLimitAndOffset(filter *TemperatureReadingFilter) (int, int) {
	page, perPage := 1, 10
	if filter != nil {
		page, perPage = pagination.ConstrainPageAndPerPage(filter.Page, filter.PerPage)
	}
	return pagination.GetLimitAndOffset(page, perPage)
}
