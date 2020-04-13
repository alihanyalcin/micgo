package {servicename}

import (
	"context"
	"{project}/internal/pkg/bootstrap/container"
	"{project}/internal/pkg/bootstrap/startup"
	"{project}/internal/pkg/di"
	"{project}/internal/pkg/logger"
	"{project}/internal/pkg/db/interfaces"
	"sync"
)

var LoggingClient logger.LoggingClient
var DbClient interfaces.DBClient

func BootstrapHandler(wg *sync.WaitGroup, ctx context.Context, startupTimer startup.Timer, dic *di.Container) bool {

	LoggingClient = container.LoggingClientFrom(dic.Get)
	LoggingClient.Info("{servicename} microservice is initializing.")

	DbClient = container.DBClientFrom(dic.Get)

	return true
}
