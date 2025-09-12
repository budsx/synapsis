package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database          Database
	Log               Log
	GRPCPort          int
	RESTPort          int
	Rabbitmq          Rabbitmq
	InventoryHost     string
	InventoryPort     int
	TopicReserveStock string
	TopicReleaseStock string
	Redis             Redis
}

type Database struct {
	DriverName string
	Host       string
	Port       int
	Username   string
	Password   string
	DbName     string
}

type Rabbitmq struct {
	RabbitmqURL string
}

type Log struct {
	Level int
}

type Redis struct {
	RedisHost     string
	RedisPassword string
	RedisDB       int
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
		GRPCPort: getEnvInt("GRPC_PORT", 8002),
		RESTPort: getEnvInt("REST_PORT", 8003),
		Rabbitmq: Rabbitmq{
			RabbitmqURL: getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		},
		InventoryHost:     getEnv("INVENTORY_HOST", "localhost"),
		InventoryPort:     getEnvInt("INVENTORY_PORT", 8000),
		TopicReserveStock: getEnv("TOPIC_RESERVE_STOCK", "reserve.stock"),
		TopicReleaseStock: getEnv("TOPIC_RELEASE_STOCK", "release.stock"),
		Redis: Redis{
			RedisHost:     getEnv("REDIS_HOST", "localhost"),
			RedisPassword: getEnv("REDIS_PASSWORD", "redis"),
			RedisDB:       getEnvInt("REDIS_DB", 0),
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
