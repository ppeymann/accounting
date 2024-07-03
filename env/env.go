package env

import "os"

// GetStringDefault returns string from environment variable for specific key or returns expected default value.
func GetStringDefault(key string, def string) string {
	val := os.Getenv(string(key))
	if len(val) == 0 {
		return def
	}

	return val
}

func IsProduction() bool {
	val := GetStringDefault("GIN_MODE", "debug")
	if val == "release" {
		return true
	}
	return false
}
