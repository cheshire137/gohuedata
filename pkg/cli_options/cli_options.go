package cli_options

import (
	"flag"
	"fmt"
)

type SensorSelection string
type LightSelection string
type TemperatureUnits string

const AllBridges = -1

const (
	AllSensors  SensorSelection = "all"
	NoSensors                   = "none"
	Motion                      = "motion"
	Temperature                 = "temperature"
)

const (
	AllLights LightSelection = "all"
	NoLights  LightSelection = "none"
)

const (
	Celsius     TemperatureUnits = "C"
	Fahrenheit                   = "F"
	Unspecified                  = ""
)

// Command-line options and flags
type CliOptions struct {
	BridgeSelection  int
	LightSelection   LightSelection
	SensorSelection  SensorSelection
	TemperatureUnits TemperatureUnits
	ConfigPath       string
	QuietMode        bool
}

func NewCliOptions(bridgeSelection int, lightSelection LightSelection, sensorSelection SensorSelection, tempUnits TemperatureUnits, configPath string, quietMode bool) *CliOptions {
	return &CliOptions{
		BridgeSelection:  bridgeSelection,
		LightSelection:   lightSelection,
		SensorSelection:  sensorSelection,
		TemperatureUnits: tempUnits,
		ConfigPath:       configPath,
		QuietMode:        quietMode,
	}
}

func ParseOptions() *CliOptions {
	var bridgeSelection int
	flag.IntVar(&bridgeSelection, "b", 0, fmt.Sprintf("Philips Hue bridge index from config.yml, starts at 1. Use %d "+
		"to hit all bridges specified in the config file.", AllBridges))

	var lightSelection string
	flag.StringVar(&lightSelection, "lights", "all", "Whether to load lights; defaults to all; choose 'all' or 'none'")

	var sensorSelection string
	flag.StringVar(&sensorSelection, "sensors", "all",
		"Which sensors to display, if any; defaults to all; choose 'all', 'motion', 'temperature', or 'none'")

	var tempUnits string
	flag.StringVar(&tempUnits, "t", "", "Temperature units to use; choose from `F` for Fahrenheit or `C` for "+
		"Celsius; defaults to the temperature_units setting in config.yml")

	var configPath string
	flag.StringVar(&configPath, "config", "config.yml", "Path to YAML configuration file; defaults to config.yml")

	var quietMode bool
	flag.BoolVar(&quietMode, "quiet", false, "Whether to reduce output; defaults to false; choose true or false.")

	flag.Parse()

	return NewCliOptions(bridgeSelection, LightSelection(lightSelection), SensorSelection(sensorSelection),
		TemperatureUnits(tempUnits), configPath, quietMode)
}

func (o *CliOptions) LoadAllBridges() bool {
	return o.BridgeSelection == AllBridges
}

func (o *CliOptions) LoadLights() bool {
	return o.LightSelection != NoLights
}

func (o *CliOptions) LoadSensors() bool {
	return o.SensorSelection != NoSensors
}

func (o *CliOptions) FahrenheitSpecified(fahrenheitFallback bool) bool {
	if o.TemperatureUnits != Unspecified {
		return o.TemperatureUnits == Fahrenheit
	}
	return fahrenheitFallback
}
