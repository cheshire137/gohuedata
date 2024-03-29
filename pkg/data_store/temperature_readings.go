package data_store

import (
	"fmt"
	"strings"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/pagination"
)

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

func (ds *DataStore) TotalDailyTemperatureReadings(filter *TemperatureReadingFilter) (int, error) {
	whereClause, values := buildTemperatureReadingWhereConditions(filter)
	queryStr := "SELECT COUNT(*) FROM temperature_readings" + temperatureReadingJoins + whereClause +
		"GROUP BY strftime('%Y-%m-%d', temperature_readings.last_updated)"
	fmt.Println(queryStr)
	var count int
	err := ds.db.QueryRow(queryStr, values...).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (ds *DataStore) LoadTemperatureReading(sensorID string, timestamp string, fahrenheit bool) (*TemperatureReading, error) {
	queryStr := `SELECT temperature_readings.last_updated AS timestamp,
			temperature_readings.temperature,
			temperature_readings.units,
			temperature_sensors.id AS sensor_id,
			temperature_sensors.name AS sensor_name,
			temperature_sensors.bridge_ip_address AS bridge_ip_address,
			hue_bridges.name AS bridge_name
		FROM temperature_readings` + temperatureReadingJoins + `
		WHERE temperature_readings.temperature_sensor_id = ?
			AND temperature_readings.last_updated = ?
		LIMIT 1`

	var reading TemperatureReading
	var sensor TemperatureSensor
	var bridge HueBridge

	err := ds.db.QueryRow(queryStr, sensorID, timestamp).Scan(&reading.Timestamp, &reading.Temperature, &reading.Units,
		&reading.temperatureSensorID, &sensor.Name, &bridge.IPAddress, &bridge.Name)
	if err != nil {
		return nil, err
	}

	reading.Temperature = hue_api.ConvertTemperature(reading.Temperature, reading.Units, fahrenheit)
	if fahrenheit {
		reading.Units = "F"
	} else {
		reading.Units = "C"
	}
	reading.ID = fmt.Sprintf("%s%s%.1f%s", reading.temperatureSensorID, reading.Timestamp, reading.Temperature,
		reading.Units)
	sensor.ID = reading.temperatureSensorID
	bridge.ID = bridge.IPAddress
	sensor.Bridge = &bridge
	reading.TemperatureSensor = &sensor

	return &reading, nil
}

func (ds *DataStore) LoadDailyTemperatureReadings(filter *TemperatureReadingFilter) ([]*TemperatureReadingSummary, error) {
	limit, offset := temperatureReadingLimitAndOffset(filter)
	whereClause, values := buildTemperatureReadingWhereConditions(filter)
	queryStr := `SELECT strftime('%Y-%m-%d', temperature_readings.last_updated) AS day,
		MIN(temperature_readings.temperature) as min_temperature,
		AVG(temperature_readings.temperature) AS avg_temperature,
		MAX(temperature_readings.temperature) AS max_temperature,
		temperature_readings.units,
		temperature_sensors.id AS sensor_id,
		temperature_sensors.name AS sensor_name,
		temperature_sensors.bridge_ip_address AS bridge_ip_address,
		hue_bridges.name AS bridge_name
		FROM temperature_readings` + temperatureReadingJoins + whereClause + `
		GROUP BY strftime('%Y-%m-%d', temperature_readings.last_updated), temperature_readings.units,
			temperature_sensors.id, temperature_sensors.name, temperature_sensors.bridge_ip_address, hue_bridges.name
		ORDER BY strftime('%Y-%m-%d', temperature_readings.last_updated) DESC, temperature_sensors.name ASC,
			hue_bridges.name ASC
		LIMIT ? OFFSET ?`
	values = append(values, limit, offset)

	rows, err := ds.db.Query(queryStr, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readingSummaries []*TemperatureReadingSummary
	sensorsByID := map[string]*TemperatureSensor{}
	bridgesByIPAddress := map[string]*HueBridge{}
	sensors := []*TemperatureSensor{}

	for rows.Next() {
		var readingSummary TemperatureReadingSummary
		var sensorName string
		var bridgeIPAddress string
		var bridgeName string

		err = rows.Scan(&readingSummary.Timestamp, &readingSummary.MinTemperature, &readingSummary.AvgTemperature,
			&readingSummary.MaxTemperature, &readingSummary.Units, &readingSummary.temperatureSensorID,
			&sensorName, &bridgeIPAddress, &bridgeName)
		if err != nil {
			return nil, err
		}

		if filter != nil {
			readingSummary.MinTemperature = hue_api.ConvertTemperature(readingSummary.MinTemperature, readingSummary.Units,
				filter.Fahrenheit)
			readingSummary.MaxTemperature = hue_api.ConvertTemperature(readingSummary.MaxTemperature, readingSummary.Units,
				filter.Fahrenheit)
			readingSummary.AvgTemperature = hue_api.ConvertTemperature(readingSummary.AvgTemperature, readingSummary.Units,
				filter.Fahrenheit)
			if filter.Fahrenheit {
				readingSummary.Units = "F"
			} else {
				readingSummary.Units = "C"
			}
		}

		sensor, ok := sensorsByID[readingSummary.temperatureSensorID]
		if !ok {
			sensor = &TemperatureSensor{
				ID:              readingSummary.temperatureSensorID,
				Name:            sensorName,
				bridgeIPAddress: bridgeIPAddress,
			}
			sensorsByID[readingSummary.temperatureSensorID] = sensor
			sensors = append(sensors, sensor)
		}

		bridge, ok := bridgesByIPAddress[bridgeIPAddress]
		if !ok {
			bridge = &HueBridge{
				ID:        bridgeIPAddress,
				IPAddress: bridgeIPAddress,
				Name:      bridgeName,
			}
			bridgesByIPAddress[bridgeIPAddress] = bridge
		}

		sensor.Bridge = bridge
		readingSummary.TemperatureSensor = sensor

		readingSummaries = append(readingSummaries, &readingSummary)
	}

	return readingSummaries, nil
}

func (ds *DataStore) LoadTemperatureReadings(filter *TemperatureReadingFilter) ([]*TemperatureReading, error) {
	limit, offset := temperatureReadingLimitAndOffset(filter)
	whereClause, values := buildTemperatureReadingWhereConditions(filter)
	queryStr := `SELECT temperature_readings.last_updated AS timestamp,
			temperature_readings.temperature,
			temperature_readings.units,
			temperature_sensors.id AS sensor_id,
			temperature_sensors.name AS sensor_name,
			temperature_sensors.bridge_ip_address AS bridge_ip_address,
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
	sensorsByID := map[string]*TemperatureSensor{}
	bridgesByIPAddress := map[string]*HueBridge{}
	sensors := []*TemperatureSensor{}

	for rows.Next() {
		var reading TemperatureReading
		var sensorName string
		var bridgeIPAddress string
		var bridgeName string

		err = rows.Scan(&reading.Timestamp, &reading.Temperature, &reading.Units, &reading.temperatureSensorID,
			&sensorName, &bridgeIPAddress, &bridgeName)
		if err != nil {
			return nil, err
		}

		reading.ID = fmt.Sprintf("%s%s%.1f%s", reading.temperatureSensorID, reading.Timestamp, reading.Temperature,
			reading.Units)

		if filter != nil {
			reading.Temperature = hue_api.ConvertTemperature(reading.Temperature, reading.Units, filter.Fahrenheit)
			if filter.Fahrenheit {
				reading.Units = "F"
			} else {
				reading.Units = "C"
			}
		}

		sensor, ok := sensorsByID[reading.temperatureSensorID]
		if !ok {
			sensor = &TemperatureSensor{
				ID:              reading.temperatureSensorID,
				Name:            sensorName,
				bridgeIPAddress: bridgeIPAddress,
			}
			sensorsByID[reading.temperatureSensorID] = sensor
			sensors = append(sensors, sensor)
		}

		bridge, ok := bridgesByIPAddress[bridgeIPAddress]
		if !ok {
			bridge = &HueBridge{
				ID:        bridgeIPAddress,
				IPAddress: bridgeIPAddress,
				Name:      bridgeName,
			}
			bridgesByIPAddress[bridgeIPAddress] = bridge
		}

		sensor.Bridge = bridge
		reading.TemperatureSensor = sensor

		readings = append(readings, &reading)
	}

	setLastUpdatedOnSensors(sensors, readings)

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

	if filter.SensorID != "" {
		conditions = append(conditions, "temperature_readings.temperature_sensor_id = ?")
		values = append(values, filter.SensorID)
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
