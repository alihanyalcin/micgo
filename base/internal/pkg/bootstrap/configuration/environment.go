package configuration

import (
	"os"
	"project/internal/pkg/config"
	"strconv"
)

const (
	envKeyStartupDuration = "startup_duration"
	envKeyStartupInterval = "startup_interval"
)

// OverrideFromEnvironment overrides the registryInfo values from an environment variable value (if it exists).
func OverrideFromEnvironment(startup config.StartupInfo) config.StartupInfo {

	//	Override the startup timer configuration, if provided.
	if env := os.Getenv(envKeyStartupDuration); env != "" {
		if n, err := strconv.ParseInt(env, 10, 0); err == nil && n > 0 {
			startup.Duration = int(n)
		}
	}
	if env := os.Getenv(envKeyStartupInterval); env != "" {
		if n, err := strconv.ParseInt(env, 10, 0); err == nil && n > 0 {
			startup.Interval = int(n)
		}
	}

	return startup
}
