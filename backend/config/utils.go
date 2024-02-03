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

	// Convert string to bool or panic
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic("Invalid value for " + key + " environment variable, must be a boolean")
	}

	return boolValue
}

func EnvInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	// Convert string to int or panic
	strconvValue, err := strconv.Atoi(value)
	if err != nil {
		panic("Invalid value for " + key + " environment variable, must be an integer")
	}

	return strconvValue
}
