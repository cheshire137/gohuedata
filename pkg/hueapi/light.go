package hueapi

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

type LightState struct {
	On               bool      `json:"on"`
	Alert            string    `json:"alert"`
	ColorMode        string    `json:"colormode"`
	Mode             string    `json:"mode"`
	Reachable        bool      `json:"reachable"`
	Brightness       int       `json:"bri"`
	Hue              int       `json:"hue"`
	Saturation       int       `json:"sat"`
	Effect           string    `json:"effect"`
	XY               []float32 `json:"xy"`
	ColorTemperature int       `json:"ct"`
}

type LightCapabilities struct {
	Certified bool         `json:"certified"`
	Control   LightControl `json:"control"`
	Streaming struct {
		Renderer bool `json:"renderer"`
		Proxy    bool `json:"proxy"`
	} `json:"streaming"`
}

type LightControl struct {
	MinDimLevel      int         `json:"mindimlevel"`
	MaxLumen         int         `json:"maxlumen"`
	ColorGamutType   string      `json:"colorgamuttype"`
	ColorGamut       [][]float32 `json:"colorgamut"`
	ColorTemperature struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"ct"`
}

type LightConfig struct {
	Archetype string `json:"archetype"`
	Function  string `json:"function"`
	Direction string `json:"direction"`
	Startup   struct {
		Mode       string `json:"mode"`
		Configured bool   `json:"configured"`
	} `json:"startup"`
}
