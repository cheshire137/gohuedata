package hue_api

// https://developers.meethue.com/develop/hue-api/groupds-api
type Group struct {
	ID        string
	Name      string      `json:"name"`
	Type      string      `json:"type"`
	LightIDs  []string    `json:"lights"`
	SensorIDs []string    `json:"sensors"`
	Action    GroupAction `json:"action"`
	State     GroupState  `json:"state"`
	Class     string      `json:"class"`
}
