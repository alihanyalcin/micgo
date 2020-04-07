package container

import (
	"{project}/internal/pkg/db/interfaces"
	"{project}/internal/pkg/di"
)

// DBClientInterfaceName contains the name of the interfaces.DBClient implementation in the DIC.
var DBClientInterfaceName = di.TypeInstanceToName((*interfaces.DBClient)(nil))

// DBClientFrom helper function queries the DIC and returns the interfaces.DBClient implementation.
func DBClientFrom(get di.Get) interfaces.DBClient {
	return get(DBClientInterfaceName).(interfaces.DBClient)
}
