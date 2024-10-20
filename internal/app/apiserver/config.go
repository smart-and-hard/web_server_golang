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

/*
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == " " {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}

// CONFIG_PATH = ./path/to/config/file.yaml ssp
func fetchConfigPath() (result string) {

	flag.StringVar(&result, "config", "", "path to config file")
	flag.Parse()

	if result == "" {
		result = os.Getenv("CONFIG_PATH")
	}

	return
}
*/
