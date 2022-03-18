package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	Port     string
	Rpc      string
	Db       string
	LogLevel log.Level
}

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Panicf("Failed to get from environment: %s", key)
	}
	return val
}

func parseLogLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case "debug":
		return log.DebugLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	case "panic":
		return log.PanicLevel
	case "info":
		return log.InfoLevel
	}
	log.Warnf("⚠️ Invalid log level: %s. Using Error as Logging Level.", level)
	return log.ErrorLevel
}

func Load() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
	port := getEnv("PORT")
	rpc := getEnv("RPC")
	db := getEnv("DB")
	logLevel := parseLogLevel(getEnv("LOG_LEVEL"))
	return AppConfig{port, rpc, db, logLevel}
}
