package config

import (
	"os"
)

// Config represents a set of configuration values.
type Config struct {
	IsDebug bool
}

// New returns a pointer to a Config struct.
func New() *Config {
	debug := os.Getenv("DEBUG")
	isDebug := false

	if debug == "true" {
		isDebug = true
	}

	return &Config{
		IsDebug: isDebug,
	}
}
