package hue_api

import (
	"fmt"
	"net/url"
)

type Bridge struct {
	Name      string `yaml:"name"`
	IPAddress string `yaml:"ip_address"`
	Username  string `yaml:"username"`
}

func (b *Bridge) GetApiUrl() (string, error) {
	if b.Username == "" {
		return "", fmt.Errorf("no username specified for bridge")
	}
	if b.IPAddress == "" {
		return "", fmt.Errorf("no IP address specified for bridge")
	}
	uri, err := url.Parse("http://" + b.IPAddress + "/api/" + b.Username)
	if err != nil {
		return "", err
	}
	return uri.String(), nil
}

func (b *Bridge) String() string {
	if b.Username != "" && b.IPAddress != "" && b.Name != "" {
		return fmt.Sprintf("%s -- %s via %s", b.Name, b.IPAddress, b.Username)
	}
	if b.Name != "" && b.IPAddress != "" {
		return fmt.Sprintf("%s -- %s (not authenticated)", b.Name, b.IPAddress)
	}
	if b.Name != "" && b.Username != "" {
		return fmt.Sprintf("%s via %s", b.Name, b.Username)
	}
	if b.IPAddress != "" && b.Username != "" {
		return fmt.Sprintf("%s via %s", b.IPAddress, b.Username)
	}
	if b.Name != "" {
		return b.Name
	}
	if b.IPAddress != "" {
		return b.IPAddress
	}
	if b.Username != "" {
		return "Unnamed bridge via " + b.Username
	}
	return "Unnamed bridge"
}

func BridgeNames(bridges []*Bridge) []string {
	names := make([]string, len(bridges))
	for i, bridge := range bridges {
		names[i] = bridge.Name
	}
	return names
}
