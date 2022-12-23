package hue_api

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
