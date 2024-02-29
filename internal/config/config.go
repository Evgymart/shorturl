package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env        string         `yaml:"env" env:"ENV" env-required:"true"`
	HttpServer HTTPServer     `yaml:"http_server" env-required:"true"`
	DbServer   DatabaseServer `yaml:"database_server" env-required:"true"`
}

type HTTPServer struct {
	Addr        string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DatabaseServer struct {
	Ip       string `yaml:"ip" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
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
