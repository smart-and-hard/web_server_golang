package store

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}

//database_url = "host=localhost dbname=restapi_dev sslmode=disable"
