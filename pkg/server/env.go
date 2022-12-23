package server

import "github.com/cheshire137/gohuedata/pkg/data_store"

type Env struct {
	ds *data_store.DataStore
}

func NewEnv(ds *data_store.DataStore) *Env {
	return &Env{ds: ds}
}
