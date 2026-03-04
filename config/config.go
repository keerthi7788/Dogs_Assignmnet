package configs

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type ApxConfig struct {
	Port         string   `koanf:"port"`
	Postgres     Postgres `koanf:"postgres"`
	PostgresTest Postgres `koanf:"postgres_test"`
}

type Postgres struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Dbname   string `koanf:"db"`
}

var DefaultConfig = ApxConfig{
	Port: ":8082",
	Postgres: Postgres{
		Host:     "localhost",
		Port:     5450,
		User:     "dogs",
		Password: "dogs123",
		Dbname:   "dogs_db",
	},
	PostgresTest: Postgres{
		Host:     "localhost",
		Port:     5450,
		User:     "dogs",
		Password: "dogs123",
		Dbname:   "dogs_db",
	},
}

// Global koanf instance
var k = koanf.New(".") // "." is the key delimiter

// Load reads the configuration from YAML and returns ApxConfig
func Load() ApxConfig {
	var cfg ApxConfig

	// Try loading from file
	if err := k.Load(file.Provider("/etc/secrets/config.yaml"), yaml.Parser()); err != nil {
		log.Println("Warning: could not load config.yaml, using default config:", err)
		return DefaultConfig
	}

	// Unmarshal into ApxConfig struct
	if err := k.Unmarshal("", &cfg); err != nil {
		log.Println("Warning: could not parse config.yaml, using default config:", err)
		return DefaultConfig
	}

	return cfg
}
