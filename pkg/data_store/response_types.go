package data_store

type TemperatureReadingsResponse struct {
	TemperatureReadings []*TemperatureReading `json:"temperatureReadings"`
	Page                int                   `json:"page"`
	TotalPages          int                   `json:"totalPages"`
	TotalCount          int                   `json:"totalCount"`
}

type TemperatureSensorsResponse struct {
	TemperatureSensors []*TemperatureSensor `json:"temperatureSensors"`
	Page               int                  `json:"page"`
	TotalPages         int                  `json:"totalPages"`
	TotalCount         int                  `json:"totalCount"`
}

type TemperatureSensorResponse struct {
	TemperatureSensor *TemperatureSensorExtended `json:"temperatureSensor"`
}

type TemperatureSensorsLiveResponse struct {
	TemperatureSensors []*TemperatureSensorExtended `json:"temperatureSensors"`
	Page               int                          `json:"page"`
	TotalPages         int                          `json:"totalPages"`
	TotalCount         int                          `json:"totalCount"`
}

type GroupsLiveResponse struct {
	Groups     []*Group `json:"groups"`
	Page       int      `json:"page"`
	TotalPages int      `json:"totalPages"`
	TotalCount int      `json:"totalCount"`
}

type TemperatureReading struct {
	ID                  string             `json:"id"`
	TemperatureSensor   *TemperatureSensor `json:"temperatureSensor"`
	Timestamp           string             `json:"timestamp"`
	Temperature         float32            `json:"temperature"`
	Units               string             `json:"units"`
	temperatureSensorID string
}

type TemperatureSensor struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Bridge          *HueBridge `json:"bridge"`
	LastUpdated     string     `json:"lastUpdated"`
	bridgeIPAddress string
}

type TemperatureSensorExtended struct {
	TemperatureSensor
	LatestReading *TemperatureReading `json:"latestReading"`
}

type HueBridge struct {
	ID        string `json:"id"`
	IPAddress string `json:"ipAddress"`
	Name      string `json:"name"`
}

type Group struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	TotalLights  int        `json:"totalLights"`
	TotalSensors int        `json:"totalSensors"`
	Bridge       *HueBridge `json:"bridge"`
	Class        string     `json:"class"`
}
