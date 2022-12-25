package data_store

type TemperatureReading struct {
	TemperatureSensor *TemperatureSensor `json:"temperatureSensor"`
	LastUpdated       string             `json:"lastUpdated"`
	Temperature       float32            `json:"temperature"`
	Units             string             `json:"units"`
}

type TemperatureSensor struct {
	Name   string     `json:"name"`
	Bridge *HueBridge `json:"bridge"`
}

type HueBridge struct {
	Name string `json:"name"`
}
