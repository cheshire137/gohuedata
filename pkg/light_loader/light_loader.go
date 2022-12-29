package light_loader

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/util"
)

type LightLoader struct {
	LightsByID map[string]hue_api.Light
}

func NewLightLoader(hueClient *hue_api.Client) (*LightLoader, error) {
	lightsByID, err := hueClient.GetLights()
	if err != nil {
		return nil, err
	}

	return &LightLoader{LightsByID: lightsByID}, nil
}

func (ll *LightLoader) TotalLights() int {
	return len(ll.LightsByID)
}

func (ll *LightLoader) SortedLightIDs() ([]string, error) {
	intIDs := make([]int, 0, len(ll.LightsByID))
	for id := range ll.LightsByID {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, err
		}
		intIDs = append(intIDs, idInt)
	}
	sort.Ints(intIDs)
	strIDs := make([]string, 0, len(intIDs))
	for _, id := range intIDs {
		strIDs = append(strIDs, strconv.Itoa(id))
	}
	return strIDs, nil
}

func (ll *LightLoader) DisplayLights(quietMode bool) error {
	count := len(ll.LightsByID)
	units := util.Pluralize(count, "light", "lights")
	util.LogSuccess("Got %d %s%s", count, units, util.LinePunctuation(quietMode))
	if !quietMode {
		sortedIDs, err := ll.SortedLightIDs()
		if err != nil {
			return err
		}
		for _, id := range sortedIDs {
			light := ll.LightsByID[id]
			fmt.Printf("%s. %s\n", id, light.String())
		}
	}
	return nil
}
