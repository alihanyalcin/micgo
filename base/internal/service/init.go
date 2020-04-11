package {servicename}

import (
	"context"
	"{project}/internal/pkg/bootstrap/container"
	"{project}/internal/pkg/bootstrap/startup"
	"{project}/internal/pkg/di"
	"{project}/internal/pkg/logger"
	"sync"
)

var LoggingClient logger.LoggingClient

func BootstrapHandler(wg *sync.WaitGroup, ctx context.Context, startupTimer startup.Timer, dic *di.Container) bool {

	LoggingClient = container.LoggingClientFrom(dic.Get)
	LoggingClient.Info("{servicename} microservice is initializing.")

	return true
}
