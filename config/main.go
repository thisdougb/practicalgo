package config

import (
	"os"
	"strconv"
)

// Config holding struct
type Config struct{}

var defaultValues = map[string]interface{}{
	"API_PORT":         "8080",        // api listens on this port
	"REDIS_HOST":       "localhost",   // redis host name
	"REDIS_PORT":       "6379",        // redis port
	"REDIS_USERNAME":   "",            // redis host name
	"REDIS_PASSWORD":   "",            // redis host name
	"REDIS_TLS":        false,         // enables TLS connection to redis
	"REDIS_KEY_PREFIX": "golocation:", // used to give scope to keys within the redis db
}

// ValueAsStr gets a string value from the env or default
func (c *Config) ValueAsStr(key string) string {

	defaultValue := defaultValues[key].(string)
	return c.getEnvVar(key, defaultValue).(string)
}

// ValueAsInt gets a string value from the env or default
func (c *Config) ValueAsInt(key string) int {

	defaultValue := defaultValues[key].(int)
	return c.getEnvVar(key, defaultValue).(int)
}

// ValueAsBool gets a string value from the env or default
func (c *Config) ValueAsBool(key string) bool {

	defaultValue := defaultValues[key].(bool)
	return c.getEnvVar(key, defaultValue).(bool)
}

// Private methods here
func (c *Config) getEnvVar(key string, fallback interface{}) interface{} {

	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}

	switch fallback.(type) {
	case string:
		return value
	case bool:
		valueAsBool, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return valueAsBool
	case int:
		valueAsInt, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return valueAsInt
	}
	return fallback
}
