package data_store

type TemperatureReadingFilter struct {
	Page          int
	PerPage       int
	BridgeName    string
	SensorID      string
	UpdatedSince  string
	UpdatedBefore string
}

type TemperatureSensorFilter struct {
	Page       int
	PerPage    int
	BridgeName string
}
