package test

import (
	"context"
	"github.com/alihanyalcin/go-microservices/internal/pkg/bootstrap/container"
	"github.com/alihanyalcin/go-microservices/internal/pkg/bootstrap/startup"
	"github.com/alihanyalcin/go-microservices/internal/pkg/di"
	"github.com/alihanyalcin/go-microservices/internal/pkg/logger"
	"sync"
)

var LoggingClient logger.LoggingClient

func BootstrapHandler(wg *sync.WaitGroup, ctx context.Context, startupTimer startup.Timer, dic *di.Container) bool {

	LoggingClient = container.LoggingClientFrom(dic.Get)

	return true
}
