package light_display

import (
	"fmt"

	"github.com/cheshire137/gohuedata/pkg/hueapi"
)

type LightDisplay struct {
	HueClient *hueapi.Client
}

func NewLightDisplay(hueClient *hueapi.Client) *LightDisplay {
	return &LightDisplay{HueClient: hueClient}
}

func (ld *LightDisplay) LoadLights() {
	lights, err := ld.HueClient.GetLights()
	if err != nil {
		fmt.Println("❌ Failed to get lights:", err)
		return
	}

	fmt.Printf("\n✅ Got %d light(s):\n", len(lights))
	for i, light := range lights {
		fmt.Printf("%d. %s\n", i+1, light.String())
	}
}
