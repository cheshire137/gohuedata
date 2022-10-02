package light_loader

import (
	"fmt"

	"github.com/cheshire137/gohuedata/pkg/hueapi"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type LightLoader struct {
	Lights []hueapi.Light
}

func NewLightLoader(hueClient *hueapi.Client) (*LightLoader, error) {
	lights, err := hueClient.GetLights()
	if err != nil {
		return nil, err
	}

	return &LightLoader{Lights: lights}, nil
}

func (ll *LightLoader) DisplayLights() {
	count := len(ll.Lights)
	units := util.Pluralize(count, "light", "lights")
	util.LogSuccess("Got %d %s:", count, units)
	for i, light := range ll.Lights {
		fmt.Printf("%d. %s\n", i+1, light.String())
	}
}
