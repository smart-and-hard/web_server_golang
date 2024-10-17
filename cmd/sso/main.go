package main

import (
	"flag"
	"log"
	"url-shortener/internal/app/apiserver"

	"github.com/BurntSushi/toml"
)

// const (
// 	envLocal = "local"
// 	envDev   = "dev"
// 	envProd  = "prod"
// )

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	// cfg := config.MustLoad()
	// fmt.Println(cfg)

	// logv := setupLogger(cfg.Env)
	// logv.Info("starting server", slog.String("env", cfg.Env))
	// logv.Debug("debug message are enabled")

	// storage, err := sqllite.New(cfg.StoragePath)
	// if err != nil {
	// 	log.Error("failed to init storage", sl.Err(err))
	// 	os.Exit(1)
	// }

	// _ = storage

}

// func setupLogger(env string) *slog.Logger {
// 	var log *slog.Logger

// 	switch env {
// 	case envLocal:
// 		log = slog.New(
// 			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
// 		)
// 	case envDev:
// 		log = slog.New(
// 			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
// 		)
// 	case envProd:
// 		log = slog.New(
// 			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
// 		)
// 	}
// 	return log
// }
