package config

import (
	"fmt"
	"os"
)

func getenvWithError(key string) (string, error) {
	if val := os.Getenv(key); val == "" {
		return "", fmt.Errorf("environmental valuable %s was not set", key)
	} else {
		return val, nil
	}
}

func getenvWithDefaultValue(key, defaultValue string) string {
	if val, err := getenvWithError(key); err != nil {
		return defaultValue
	} else {
		return val
	}
}
