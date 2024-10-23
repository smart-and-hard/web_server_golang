package apiserver

import (
	"url-shortener/internal/store"
)

type Config struct {
	BinAddr  string `toml:"BinAddr"`
	logLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  "5050",
		logLevel: "debug",
		Store:    store.NewConfig(),
	}
}
