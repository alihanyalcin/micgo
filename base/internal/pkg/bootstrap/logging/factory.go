package logging

import (
	"{project}/internal/pkg/bootstrap/interfaces"
	"{project}/internal/pkg/logger"
)

// FactoryToStdout returns a logger.LoggingClient that outputs to stdout.
func FactoryToStdout(serviceKey string) logger.LoggingClient {
	return logger.NewClientStdOut(serviceKey, logger.DebugLog)
}

// FactoryFromConfiguration returns a logger.LoggingClient based on configuration settings.
func FactoryFromConfiguration(serviceKey string, config interfaces.Configuration) logger.LoggingClient {
	var target string
	bootstrapConfig := config.GetBootstrap()
	
	target = bootstrapConfig.Logging.File
	
	return logger.NewClient(serviceKey, target, config.GetLogLevel())
}
