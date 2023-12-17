package data_store

type TemperatureReadingsResponse struct {
	TemperatureReadings []*TemperatureReading `json:"temperatureReadings"`
	Page                int                   `json:"page"`
	TotalPages          int                   `json:"totalPages"`
	TotalCount          int                   `json:"totalCount"`
}

type TemperatureReadingSummariesResponse struct {
	TemperatureReadingSummaries []*TemperatureReadingSummary `json:"temperatureReadingSummaries"`
	Page                        int                          `json:"page"`
	TotalPages                  int                          `json:"totalPages"`
	TotalCount                  int                          `json:"totalCount"`
}

type TemperatureSensorsResponse struct {
	TemperatureSensors []*TemperatureSensor `json:"temperatureSensors"`
	Page               int                  `json:"page"`
	TotalPages         int                  `json:"totalPages"`
	TotalCount         int                  `json:"totalCount"`
}

type TemperatureSensorResponse struct {
	TemperatureSensor *TemperatureSensorExtended `json:"temperatureSensor"`
	MaxTemperature    *float32                   `json:"maxTemperature"`
	MinTemperature    *float32                   `json:"minTemperature"`
	AvgTemperature    *float32                   `json:"avgTemperature"`
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

type TemperatureReadingSummary struct {
	TemperatureSensor   *TemperatureSensor `json:"temperatureSensor"`
	Timestamp           string             `json:"timestamp"`
	MinTemperature      float32            `json:"minTemperature"`
	AvgTemperature      float32            `json:"avgTemperature"`
	MaxTemperature      float32            `json:"maxTemperature"`
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
	UniqueID     string     `json:"uniqueID"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	TotalLights  int        `json:"totalLights"`
	TotalSensors int        `json:"totalSensors"`
	Bridge       *HueBridge `json:"bridge"`
	Class        string     `json:"class"`
	Lights       []*Light   `json:"lights"`
}

type GroupExtended struct {
	ID           string           `json:"id"`
	UniqueID     string           `json:"uniqueID"`
	Name         string           `json:"name"`
	Type         string           `json:"type"`
	TotalLights  int              `json:"totalLights"`
	TotalSensors int              `json:"totalSensors"`
	Bridge       *HueBridge       `json:"bridge"`
	Class        string           `json:"class"`
	Lights       []*LightExtended `json:"lights"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Light struct {
	UniqueID string     `json:"uniqueID"`
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Bridge   *HueBridge `json:"bridge"`
}

type LightExtended struct {
	Light
	LatestState *LightState `json:"latestState"`
}

type LightState struct {
	Timestamp string `json:"timestamp"`
	On        bool   `json:"on"`
	Light     *Light `json:"light"`
}
