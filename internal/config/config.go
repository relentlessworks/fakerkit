package config

import (
	"flag"
	"os"
)

// Config holds all application configuration.
type Config struct {
	Port string
	Host string
}

// Default returns the default configuration.
func Default() *Config {
	return &Config{
		Port: "7118",
		Host: "0.0.0.0",
	}
}

// FromFlags parses command-line flags and environment variables.
// Precedence: defaults < env < flags.
func FromFlags() *Config {
	c := Default()

	// Environment variables
	if p := os.Getenv("FAKERKIT_PORT"); p != "" {
		c.Port = p
	}
	if h := os.Getenv("FAKERKIT_HOST"); h != "" {
		c.Host = h
	}

	// Flags (override env)
	flag.StringVar(&c.Port, "port", c.Port, "HTTP listen port")
	flag.StringVar(&c.Host, "host", c.Host, "HTTP listen host")
	flag.Parse()

	return c
}

// Addr returns the full listen address.
func (c *Config) Addr() string {
	return c.Host + ":" + c.Port
}
