package config

import (
	"{project}/internal/pkg/bootstrap/interfaces"
	"{project}/internal/pkg/config"
)

type ConfigurationStruct struct {
	Service   config.ServiceInfo
	Logging   config.LoggingInfo
	Startup   config.StartupInfo
	Databases config.DatabaseInfo
}

// GetBootstrap returns the configuration elements required by the bootstrap.
func (c *ConfigurationStruct) GetBootstrap() interfaces.BootstrapConfiguration {
	return interfaces.BootstrapConfiguration{
		Service: c.Service,
		Logging: c.Logging,
		Startup: c.Startup,
	}
}

// GetLogLevel returns the current ConfigurationStruct's log level.
func (c *ConfigurationStruct) GetLogLevel() string {
	return c.Logging.LogLevel
}

// GetDatabaseInfo returns a database information.
func (c *ConfigurationStruct) GetDatabaseInfo() config.DatabaseInfo {
	return c.Databases
}
