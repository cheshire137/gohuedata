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

func (b *Bridge) GetUrl() (*url.URL, error) {
	return url.Parse("http://" + b.IPAddress)
}

func (b *Bridge) String() string {
	if b.Username != "" {
		return fmt.Sprintf("%s -- %s via %s", b.Name, b.IPAddress, b.Username)
	}
	return fmt.Sprintf("%s -- %s", b.Name, b.IPAddress)
}
