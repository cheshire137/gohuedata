package hue_api

type SensorInput struct {
	RepeatIntervals []int         `json:"repeatintervals"`
	Events          []SensorEvent `json:"events"`
}
