package hueapi

import "net/url"

type Bridge struct {
	Name      string `yaml:"name"`
	IPAddress string `yaml:"ip_address"`
	Username  string `yaml:"username"`
}

func (b *Bridge) GetUrl() (*url.URL, error) {
	return url.Parse("http://" + b.IPAddress)
}
