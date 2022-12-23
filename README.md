# gohuedata

Command-line tool written in Go to get data from your Philips Hue system. Just a little work-in-progress that I'm fiddling with.

## How to use/develop

1. `cp config.yml.example config.yml`

    You can move the config.yml file wherever you want or rename it, but you'll need to pass the `-config` option and specify its path if you do.

1. Update config.yml with details about your Philips Hue bridges. See [Get started](https://developers.meethue.com/develop/get-started-2/) docs from Philips about how to get your bridge IP address and a username.
1. `go run cmd/gohuedata/main.go`

    You will be prompted to select a bridge if your config file specifies more than one. By default, all lights and
    sensors on the selected bridge will be shown.

### Options

- **`-b`** - Specify a bridge via the index (starting at 1) of the bridge in your config file. For example, to specify
the first bridge in your config:

    `go run cmd/gohuedata/main.go -b 1`

- **`-lights`** - Whether to fetch lights on the chosen bridge. Choose between `all` and `none`. Defaults to `all`.
Example:

    `go run cmd/gohuedata/main.go -lights none`

- **`-sensors`** - Which sensors to display from the chosen bridge, if any. Choose between `all`, `temperature`,
`motion`, and `none`. Defaults to `all`. Example:

    `go run cmd/gohuedata/main.go -sensors temperature`

- **`-t`** - Which units to use for temperature display, to override the `temperature_units` setting in config file.
Choose between `F` for Fahrenheit and `C` for Celsius. Defaults to the config file setting. Example:

    `go run cmd/gohuedata/main.go -t C`

- **`-config`** - Specify the path to the YAML configuration file.
