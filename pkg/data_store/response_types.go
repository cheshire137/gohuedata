package data_store

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

type HueBridge struct {
	ID        string `json:"id"`
	IPAddress string `json:"ipAddress"`
	Name      string `json:"name"`
}
