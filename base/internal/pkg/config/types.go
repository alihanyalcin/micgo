package config

import (
	"fmt"
)

// ServiceInfo contains configuration settings necessary for the basic operation of any EdgeX service.
type ServiceInfo struct {
	// BootTimeout indicates, in milliseconds, how long the service will retry connecting to upstream dependencies
	// before giving up. Default is 30,000.
	BootTimeout int
	// Host is the hostname or IP address of the service.
	Host string
	// Port is the HTTP port of the service.
	Port int
	// The protocol that should be used to call this service
	Protocol string
	// StartupMsg specifies a string to log once service
	// initialization and startup is completed.
	StartupMsg string
	// Timeout specifies a timeout (in milliseconds) for
	// processing REST calls from other services.
	Timeout int
}

// Url provides a way to obtain the full url of the host service for use in initialization or, in some cases,
// responses to a caller.
func (s ServiceInfo) Url() string {
	url := fmt.Sprintf("%s://%s:%v", s.Protocol, s.Host, s.Port)
	return url
}

// LoggingInfo provides basic parameters related to where logs should be written.
type LoggingInfo struct {
	EnableRemote bool
	File         string
	LogLevel     string
}

// StartupInfo provides the startup timer values which are applied to the StartupTimer created at boot.
type StartupInfo struct {
	Duration int
	Interval int
}

// DatabaseInfo defines the parameters necessary for connecting to the desired persistence layer.
type DatabaseInfo Database

type Database struct {
	Credentials
	Type    string
	Timeout int
	Host    string
	Port    int
	Name    string
}

// Credentials encapsulates username-password attributes.
type Credentials struct {
	Username string
	Password string
}
