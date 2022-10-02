package light_display

import (
	"fmt"

	"github.com/cheshire137/gohuedata/pkg/hueapi"
)

type LightDisplay struct {
	Lights []hueapi.Light
}

func NewLightDisplay(hueClient *hueapi.Client) (*LightDisplay, error) {
	lights, err := hueClient.GetLights()
	if err != nil {
		return nil, err
	}

	return &LightDisplay{Lights: lights}, nil
}

func (ld *LightDisplay) DisplayLights() {
	fmt.Printf("\n✅ Got %d light(s):\n", len(ld.Lights))
	for i, light := range ld.Lights {
		fmt.Printf("%d. %s\n", i+1, light.String())
	}
}
