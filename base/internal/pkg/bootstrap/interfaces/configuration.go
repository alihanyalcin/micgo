package interfaces

import (
	"{project}/internal/pkg/config"
)

// BootstrapConfiguration defines the configuration elements required by the bootstrap.
type BootstrapConfiguration struct {
	Service config.ServiceInfo
	Logging config.LoggingInfo
	Startup config.StartupInfo
}

// Configuration interface provides an abstraction around a configuration struct.
type Configuration interface {
	// GetBootstrap returns the configuration elements required by the bootstrap.
	GetBootstrap() BootstrapConfiguration

	// GetLogLevel returns the current ConfigurationStruct's log level.
	GetLogLevel() string
}
