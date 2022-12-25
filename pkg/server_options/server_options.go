package server_options

import (
	"flag"
	"fmt"
)

type ServerOptions struct {
	ConfigPath   string
	Port         int
	FrontendPort int
}

func NewServerOptions(configPath string, port int, frontendPort int) *ServerOptions {
	return &ServerOptions{ConfigPath: configPath, Port: port, FrontendPort: frontendPort}
}

// Keep in sync with 'proxy' port in ui/package.json and `REACT_APP_BACKEND_PORT` in ui/.env:
const DefaultPort = 8080

// Keep in sync with `PORT` environment variable in ui/.env:
const DefaultFrontendPort = 4000

func ParseOptions() *ServerOptions {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yml", "Path to YAML configuration file; defaults to config.yml")

	var port int
	flag.IntVar(&port, "p", DefaultPort, fmt.Sprintf("Port to run the server on. Defaults to %d.", DefaultPort))

	var frontendPort int
	flag.IntVar(&frontendPort, "fp", DefaultFrontendPort, fmt.Sprintf("Port the frontend runs on. Necessary to "+
		"allow API requests from the frontend to the server. Defaults to %d.", DefaultFrontendPort))

	flag.Parse()

	return NewServerOptions(configPath, port, frontendPort)
}

func (o *ServerOptions) Address() string {
	return fmt.Sprintf(":%d", o.Port)
}
