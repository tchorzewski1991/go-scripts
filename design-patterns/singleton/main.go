package main

import (
	"sync"
	"fmt"
)

// Config represents struct that we want to be initialized only once.
// It will be our singleton.
type Config struct {
	appName    string
	appVersion string
}

// SetAppName() is setter for Config appName attribute
func (c *Config) SetAppName(name string) {
	c.appName = name
}

// SetAppVersion() is setter for Config appVersion attribute
func (c *Config) SetAppVersion(version string) {
	c.appVersion = version
}

var(
	once     sync.Once
	instance *Config
)

// GetConfig() is a function where most of work related to singleton
// initialization happen. Its main responsibility is to return pointer
// to our singleton and ensures it will be initialized only once.
// This is assured due to usage of sync.Once built-in function.
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{"MyApp", "1.0.0"}
	})
	return instance
}

func main() {
	conf_1 := GetConfig()
	fmt.Printf("Pointer value of conf_1 (%v) is %p \n", *conf_1, conf_1)

	conf_2 := GetConfig()
	fmt.Printf("Pointer value of conf_2 (%v) is %p \n", *conf_2, conf_2)

	conf_2.SetAppName("NewAppName")
	conf_2.SetAppVersion("1.0.1")

	fmt.Printf("Pointer value of conf_2 (%v) is %p \n", *conf_2, conf_2)
}

