package hueapi

type SensorInput struct {
	RepeatIntervals []int         `json:"repeatintervals"`
	Events          []SensorEvent `json:"events"`
}
