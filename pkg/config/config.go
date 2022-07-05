package config

import (
	"os"

	"github.com/cheshire137/gohuedata/pkg/hueapi"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Bridges []hueapi.Bridge `yaml:"bridges"`
}

func NewConfig(path string) (*Config, error) {
	var config Config
	err := ReadConfig(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func ReadConfig(path string, config *Config) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	return decoder.Decode(config)
}
