package hueapi

type HueApi struct {
	Bridge *Bridge
}

func NewHueApi(bridge *Bridge) *HueApi {
	return &HueApi{Bridge: bridge}
}
