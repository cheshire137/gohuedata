package server

import (
	"github.com/cheshire137/gohuedata/pkg/data_store"
	options "github.com/cheshire137/gohuedata/pkg/server_options"
)

type Env struct {
	ds      *data_store.DataStore
	options *options.ServerOptions
}

func NewEnv(ds *data_store.DataStore, options *options.ServerOptions) *Env {
	return &Env{ds: ds, options: options}
}
