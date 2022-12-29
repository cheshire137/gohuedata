package group_loader

import "github.com/cheshire137/gohuedata/pkg/hue_api"

type GroupLoader struct {
	GroupsByID map[string]hue_api.Group
}

func NewGroupLoader(hueClient *hue_api.Client) (*GroupLoader, error) {
	groupsByID, err := hueClient.GetGroups()
	if err != nil {
		return nil, err
	}

	return &GroupLoader{GroupsByID: groupsByID}, nil
}

func (gl *GroupLoader) TotalGroups() int {
	return len(gl.GroupsByID)
}
