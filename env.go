package env

import (
	"os"
	"strconv"
)

// String returns the value of the environment variable named by the key
// panics if the environment variable is empty and no defaultValue is provided
//
// # If the environment variable is empty and no defaultValue is provided, it panics
//
// Example:
//
//	host := env.String("HOST", "localhost")
//	host := env.String("HOST")
func String(key string, defaultValue ...string) string {
	if value, ok := os.LookupEnv(key); ok {
		if value != "" {
			return value
		}
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	panic("Missing env variable: '" + key + "'")
}

// StringPtr returns the pointer value of the environment variable named by the key
// difference between StringPtr and String is that StringPtr returns a pointer to the value
// so that it can be nil if the environment variable is empty insteda of throwing a panic
//
// # If the environment variable is empty and no defaultValue is provided, it panics
//
// Example:
//
//	host := env.String("HOST")
//	if host != nil {
//		fmt.Println(*host)
//	}
func StringPtr(key string, defaultValue ...string) *string {
	if value, ok := os.LookupEnv(key); ok {
		if value != "" {
			return &value
		}
	}

	if len(defaultValue) > 0 {
		return &defaultValue[0]
	}

	return nil
}

// Int64 returns the value of the environment variable named by the key
// panics if the environment variable is empty and no defaultValue is provided
//
// # If the environment variable is empty and no defaultValue is provided, it panics
//
// Example:
//
//	port := env.String("PORT", "8080")
//	port := env.String("PORT")
func Int64(key string, defaultValue ...int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		if i, ok := strconv.ParseInt(value, 10, 64); ok == nil {
			return i
		}

		panic("Invalid value for env variable: '" + key + "'")
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	panic("Missing env variable: '" + key + "'")
}

// Bool returns the value of the environment variable named by the key
// default value is false if not provided
//
// # If the environment variable is empty and no defaultValue is provided, it panics
//
// Example:
//
//	isDev := env.Bool("DEBUG", true)
//	isDev := env.Bool("DEBUG") // DEBUG=1 prints true
func Bool(key string, defaultValue ...bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if i, ok := strconv.ParseBool(value); ok == nil {
			return i
		}

		panic("Invalid value for env variable: '" + key + "'")
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return false
}
