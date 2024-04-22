package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

//parse config

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8082"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-required:"true"`
	Password    string        `yaml:"password" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func MustLoadEnv() *Config {
	config := os.Getenv("CONFIG_PATH")
	if config == "" {
		log.Fatal("config is not set")
	}

	if _, err := os.Stat(config); os.IsNotExist(err) {
		log.Fatalf("config does not exist: %s", config)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(config, &cfg); err != nil {
		log.Fatalf("cannnot read config: %s", err)
	}

	return &cfg
}
