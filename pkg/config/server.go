package config

import (
	"golang-demo/domain"
	"os"
	"strconv"
	"time"
)

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func convertFloat(env string) float32 {
	v, _ := strconv.ParseFloat(os.Getenv(env), 32)
	return float32(v)
}

func ConvertEnvInt(key string) int {
	v, err := strconv.Atoi(key)
	if err != nil || key == "" {
		return 0
	}
	return v
}

func ConvertBool(env string) bool {
	v, _ := strconv.ParseBool(os.Getenv(env))
	return v
}

func DefaultValue(env string, defaultValue string) string {
	val := os.Getenv(env)
	if val == "" {
		return defaultValue
	}
	return val
}

func DefaultValueInt(env string, defaultValue int) int {
	val := ConvertInt(env)
	if val == 0 {
		return defaultValue
	}
	return val
}

func DefaultValueDuration(env string, defaultValue string) time.Duration {
	value := os.Getenv(env)
	if value == "" {
		value = defaultValue
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		panic(err)
	}
	return d
}

func DefaultValueFloat(env string, defaultValue float32) float32 {
	val := convertFloat(env)
	if val == 0 {
		return defaultValue
	}
	return val
}

func DatabasePGSQL() domain.DatabaseConfig {
	schema := "public"
	cfgSchema := os.Getenv("DB_SCHEMA")
	if cfgSchema != "" {
		schema = cfgSchema
	}

	return domain.DatabaseConfig{
		User:              os.Getenv("DB_USERNAME"),
		Password:          os.Getenv("DB_PASSWORD"),
		Database:          os.Getenv("DB_NAME"),
		Host:              os.Getenv("DB_HOST"),
		Port:              os.Getenv("DB_PORT"),
		TimeZone:          os.Getenv("DB_TIMEZONE"),
		MaxConnectionIdle: ConvertInt("DB_MAX_CON_IDLE"),
		MaxConnectionOpen: ConvertInt("DB_MAX_CON_OPEN"),
		Schema:            schema,
		Debug:             ConvertBool("DEBUG"),
		TimeOut:           time.Duration(ConvertInt("APP_TIMEOUT")) * time.Second,
	}
}
