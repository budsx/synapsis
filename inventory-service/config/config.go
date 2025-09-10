package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database Database
	Log      Log
	Rabbitmq Rabbitmq
	GRPCPort int
	RESTPort int
}

type Database struct {
	DriverName string
	Host       string
	Port       int
	Username   string
	Password   string
	DbName     string
}

type Log struct {
	Level int
}

type Rabbitmq struct {
	RabbitmqURL string
}

func Load() *Config {
	godotenv.Load()
	return &Config{
		Database: Database{
			DriverName: getEnv("DB_DRIVER", "postgres"),
			Host:       getEnv("DB_HOST", "localhost"),
			Port:       getEnvInt("DB_PORT", 5432),
			Username:   getEnv("DB_USER", "postgres"),
			Password:   getEnv("DB_PASSWORD", "postgres"),
			DbName:     getEnv("DB_NAME", "postgres"),
		},
		Log: Log{
			Level: getEnvInt("LOG_LEVEL", -1),
		},
		GRPCPort: getEnvInt("GRPC_PORT", 8000),
		RESTPort: getEnvInt("REST_PORT", 8000),
		Rabbitmq: Rabbitmq{
			RabbitmqURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return result
}

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value == "true"
}
