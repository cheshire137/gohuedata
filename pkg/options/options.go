package options

import "flag"

type SensorSelection string
type LightSelection string
type TemperatureUnits string

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

type Options struct {
	BridgeSelection  int
	LightSelection   LightSelection
	SensorSelection  SensorSelection
	TemperatureUnits TemperatureUnits
}

func NewOptions(bridgeSelection int, lightSelection LightSelection, sensorSelection SensorSelection, tempUnits TemperatureUnits) *Options {
	return &Options{
		BridgeSelection:  bridgeSelection,
		LightSelection:   lightSelection,
		SensorSelection:  sensorSelection,
		TemperatureUnits: tempUnits,
	}
}

func ParseOptions() *Options {
	var bridgeSelection int
	flag.IntVar(&bridgeSelection, "b", 0, "Philips Hue bridge index from config.yml, starts at 1")

	var lightSelection string
	flag.StringVar(&lightSelection, "lights", "all", "Whether to load lights; defaults to all; choose 'all' or 'none'")

	var sensorSelection string
	flag.StringVar(&sensorSelection, "sensors", "all",
		"Which sensors to display, if any; defaults to all; choose 'all', 'motion', 'temperature', or 'none'")

	var tempUnits string
	flag.StringVar(&tempUnits, "t", "", "Temperature units to use; choose from `F` for Fahrenheit or `C` for "+
		"Celsius; defaults to the temperature_units setting in config.yml")

	flag.Parse()

	return NewOptions(bridgeSelection, LightSelection(lightSelection), SensorSelection(sensorSelection),
		TemperatureUnits(tempUnits))
}

func (o *Options) LoadLights() bool {
	return o.LightSelection != NoLights
}

func (o *Options) LoadSensors() bool {
	return o.SensorSelection != NoSensors
}

func (o *Options) OnlyTemperatureSensors() bool {
	return o.SensorSelection == Temperature
}

func (o *Options) OnlyMotionSensors() bool {
	return o.SensorSelection == Motion
}

func (o *Options) LoadTemperatureSensors() bool {
	return o.LoadSensors() && !o.OnlyMotionSensors()
}

func (o *Options) LoadMotionSensors() bool {
	return o.LoadSensors() && !o.OnlyTemperatureSensors()
}

func (o *Options) LoadAllSensors() bool {
	return o.LoadSensors() && !o.OnlyMotionSensors() && !o.OnlyTemperatureSensors()
}

func (o *Options) AnyTemperatureUnitsSpecified() bool {
	return o.TemperatureUnits != Unspecified
}

func (o *Options) FahrenheitSpecified() bool {
	return o.TemperatureUnits == Fahrenheit
}
