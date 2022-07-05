package hueapi

type Bridge struct {
	Name      string `yaml:"name"`
	IPAddress string `yaml:"ip_address"`
	Username  string `yaml:"username"`
}
