package server_options

import (
	"flag"
	"fmt"
)

type ServerOptions struct {
	DatabasePath string
	Port         int
}

func NewServerOptions(databasePath string, port int) *ServerOptions {
	return &ServerOptions{
		DatabasePath: databasePath,
		Port:         port,
	}
}

func ParseOptions() *ServerOptions {
	var databasePath string
	flag.StringVar(&databasePath, "db", "gohuedata.db", "Path to the SQLite database file containing your "+
		"Philips Hue data, as generated by gohuedata. Defaults to gohuedata.db.")

	var port int
	flag.IntVar(&port, "p", 8080, "Port to run the server on. Defaults to 8080.")

	flag.Parse()

	return NewServerOptions(databasePath, port)
}

func (o *ServerOptions) Address() string {
	return fmt.Sprintf(":%d", o.Port)
}