package hue_api

type LightConfig struct {
	Archetype string `json:"archetype"`
	Function  string `json:"function"`
	Direction string `json:"direction"`
	Startup   struct {
		Mode       string `json:"mode"`
		Configured bool   `json:"configured"`
	} `json:"startup"`
}
