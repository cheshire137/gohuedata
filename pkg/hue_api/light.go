package hue_api

import "fmt"

// https://developers.meethue.com/develop/hue-api/lights-api/
type Light struct {
	State          LightState `json:"state"`
	SoftwareUpdate struct {
		State       string `json:"state"`
		LastInstall string `json:"lastinstall"`
	} `json:"swupdate"`
	Type             string            `json:"type"`
	Name             string            `json:"name"`
	ModelID          string            `json:"modelid"`
	ManufacturerName string            `json:"manufacturername"`
	ProductName      string            `json:"productname"`
	Capabilities     LightCapabilities `json:"capabilities"`
	Config           LightConfig       `json:"config"`
	UniqueID         string            `json:"uniqueid"`
	SoftwareVersion  string            `json:"swversion"`
	SoftwareConfigID string            `json:"swconfigid"`
	ProductID        string            `json:"productid"`
}

func (l *Light) String() string {
	onEmoji := ""
	if l.State.On {
		onEmoji = " 💡"
	}
	return fmt.Sprintf("%s%s -- %s, %s (%s)", l.Name, onEmoji, l.Type, l.ModelID, l.UniqueID)
}
