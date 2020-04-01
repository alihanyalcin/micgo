package container

import (
	"project/internal/pkg/di"
	"project/internal/pkg/logger"
)

// LoggingClientInterfaceName contains the name of the logger.LoggingClient implementation in the DIC.
var LoggingClientInterfaceName = di.TypeInstanceToName((*logger.LoggingClient)(nil))

// LoggingClientFrom helper function queries the DIC and returns the logger.loggingClient implementation.
func LoggingClientFrom(get di.Get) logger.LoggingClient {
	return get(LoggingClientInterfaceName).(logger.LoggingClient)
}
