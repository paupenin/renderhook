package config

import (
	"os"
	"strconv"
)

func Env(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func EnvBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value == "true"
}

func EnvInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	// Convert string to int (if not possible, return default value)
	strconvValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return strconvValue
}
