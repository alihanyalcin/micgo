package config

import (
	"github.com/alihanyalcin/go-microservices/internal/pkg/bootstrap/interfaces"
	"github.com/alihanyalcin/go-microservices/internal/pkg/config"
)

type ConfigurationStruct struct {
	Service   config.ServiceInfo
	Logging   config.LoggingInfo
	Startup   config.StartupInfo
	Databases config.DatabaseInfo
}

// GetBootstrap returns the configuration elements required by the bootstrap.  Currently, a copy of the configuration
// data is returned.  This is intended to be temporary -- since ConfigurationStruct drives the configuration.toml's
// structure -- until we can make backwards-breaking configuration.toml changes (which would consolidate these fields
// into an interfaces.BootstrapConfiguration struct contained within ConfigurationStruct).
func (c *ConfigurationStruct) GetBootstrap() interfaces.BootstrapConfiguration {
	// temporary until we can make backwards-breaking configuration.toml change
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

// GetDatabaseInfo returns a database information map.
func (c *ConfigurationStruct) GetDatabaseInfo() config.DatabaseInfo {
	return c.Databases
}
