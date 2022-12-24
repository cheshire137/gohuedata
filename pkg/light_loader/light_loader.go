package light_loader

import (
	"fmt"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type LightLoader struct {
	Lights []hue_api.Light
}

func NewLightLoader(hueClient *hue_api.Client) (*LightLoader, error) {
	lights, err := hueClient.GetLights()
	if err != nil {
		return nil, err
	}

	return &LightLoader{Lights: lights}, nil
}

func (ll *LightLoader) DisplayLights(quietMode bool) {
	count := len(ll.Lights)
	units := util.Pluralize(count, "light", "lights")
	util.LogSuccess("Got %d %s%s", count, units, util.LinePunctuation(quietMode))
	if !quietMode {
		for i, light := range ll.Lights {
			fmt.Printf("%d. %s\n", i+1, light.String())
		}
	}
}
