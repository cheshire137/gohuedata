package hueapi

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
		return "", fmt.Errorf("No username specified for bridge")
	}
	if b.IPAddress == "" {
		return "", fmt.Errorf("No IP address specified for bridge")
	}
	uri, err := url.Parse("http://" + b.IPAddress + "/api/" + b.Username)
	if err != nil {
		return "", err
	}
	return uri.String(), nil
}

func (b *Bridge) String() string {
	if b.Username != "" {
		return fmt.Sprintf("%s -- %s via %s", b.Name, b.IPAddress, b.Username)
	}
	return fmt.Sprintf("%s -- %s", b.Name, b.IPAddress)
}
