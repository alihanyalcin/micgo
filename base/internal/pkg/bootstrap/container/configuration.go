package container

import (
	"project/internal/pkg/bootstrap/interfaces"
	"project/internal/pkg/di"
)

// ConfigurationInterfaceName contains the name of the interfaces.Configuration implementation in the DIC.
var ConfigurationInterfaceName = di.TypeInstanceToName((*interfaces.Configuration)(nil))

// ConfigurationFrom helper function queries the DIC and returns the interfaces.Configuration implementation.
func ConfigurationFrom(get di.Get) interfaces.Configuration {
	return get(ConfigurationInterfaceName).(interfaces.Configuration)
}
