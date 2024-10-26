package apiserver

type Config struct {
	BindAddr    string `toml:"BindAddr"`
	logLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "5050",
		logLevel: "debug",
	}
}
