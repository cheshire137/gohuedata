package data_store

type TemperatureReading struct {
	TemperatureSensor   *TemperatureSensor `json:"temperatureSensor"`
	LastUpdated         string             `json:"lastUpdated"`
	Temperature         float32            `json:"temperature"`
	Units               string             `json:"units"`
	temperatureSensorID string
}

type TemperatureSensor struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Bridge          *HueBridge `json:"bridge"`
	bridgeIPAddress string
}

type HueBridge struct {
	IPAddress string `json:"ipAddress"`
	Name      string `json:"name"`
}
