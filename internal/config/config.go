package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env    string     `yaml:"env" env:"ENV" env-required:"true"`
	Server HTTPServer `yaml:"http_server" env-required:"true"`
}

type HTTPServer struct {
	Addr        string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		panic("Config not loaded, check env CONFIG_PATH")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Could not get config from: " + configPath)
	}

	return cfg
}
