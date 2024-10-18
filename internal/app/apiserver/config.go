package apiserver

import (
	"time"
)

type Config struct {
	//Env         string `yaml:"env" env-default:"local"`
	BinAddr     string `tomal:"BinAddr"`
	logLevel    string `toml:"log_level"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:5050"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  "5050",
		logLevel: "debug",
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
