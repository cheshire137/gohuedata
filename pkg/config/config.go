package config

import (
	"os"

	"github.com/cheshire137/gohuedata/pkg/hue_api"
	"github.com/cheshire137/gohuedata/pkg/options"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Bridges          []hue_api.Bridge `yaml:"bridges"`
	TemperatureUnits string           `yaml:"temperature_units"`
	DatabaseFile     string           `yaml:"database_file"`
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

func (c *Config) FahrenheitSpecified() bool {
	return c.TemperatureUnits == options.Fahrenheit
}
