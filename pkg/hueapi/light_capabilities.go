package hueapi

type LightCapabilities struct {
	Certified bool         `json:"certified"`
	Control   LightControl `json:"control"`
	Streaming struct {
		Renderer bool `json:"renderer"`
		Proxy    bool `json:"proxy"`
	} `json:"streaming"`
}
