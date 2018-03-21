package main

import (
	"fmt"
	"os"
)

// Config - Configuration Information Struct
type Config struct {
	Port string
	Addr string
}

func getenvOrElse(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

// EnvConfig - Create a new Config object from Environment
// Configuration Variables.
func EnvConfig() Config {
	port := getenvOrElse("PORT", "3000")
	addr := getenvOrElse("ADDR", "127.0.0.1")

	return Config{
		Port: port,
		Addr: addr,
	}
}

// AddressString - Get formatted URL string from supplied
// configuration information.
func (c Config) AddressString() string {
	return fmt.Sprintf("%s:%s", c.Addr, c.Port)
}
