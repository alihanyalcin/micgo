package bootstrap

import (
	"context"
	"os"
	"os/signal"
	"{project}/internal/pkg/bootstrap/configuration"
	"{project}/internal/pkg/bootstrap/container"
	"{project}/internal/pkg/bootstrap/interfaces"
	"{project}/internal/pkg/bootstrap/logging"
	"{project}/internal/pkg/bootstrap/startup"
	"{project}/internal/pkg/di"
	"{project}/internal/pkg/logger"
	"sync"
	"syscall"
)

// fatalError logs an error and exits the application.  It's intended to be used only within the bootstrap prior to
// any go routines being spawned.
func fatalError(err error, loggingClient logger.LoggingClient) {
	loggingClient.Error(err.Error())
	os.Exit(1)
}

func translateInterruptToCancel(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		signalStream := make(chan os.Signal)
		defer func() {
			signal.Stop(signalStream)
			close(signalStream)
		}()
		signal.Notify(signalStream, os.Interrupt, syscall.SIGTERM)
		select {
		case <-signalStream:
			cancel()
			return
		case <-ctx.Done():
			return
		}
	}()
}

func Run(
	configDir, profileDir, configFileName string,
	serviceKey string,
	config interfaces.Configuration,
	startupTimer startup.Timer,
	dic *di.Container,
	handlers []interfaces.BootstrapHandler) {

	loggingClient := logging.FactoryToStdout(serviceKey)

	var err error
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	translateInterruptToCancel(&wg, ctx, cancel)

	// load configuration from file
	if err = configuration.LoadFromFile(configDir, profileDir, configFileName, config); err != nil {
		fatalError(err, loggingClient)
	}

	loggingClient = logging.FactoryFromConfiguration(serviceKey, config)

	bootstrapConfig := config.GetBootstrap()

	//	Update the startup timer to reflect whatever configuration read, if anything available.
	startupTimer.UpdateTimer(bootstrapConfig.Startup.Duration, bootstrapConfig.Startup.Interval)

	dic.Update(di.ServiceConstructorMap{
		container.ConfigurationInterfaceName: func(get di.Get) interface{} {
			return config
		},
		container.LoggingClientInterfaceName: func(get di.Get) interface{} {
			return loggingClient
		},
	})

	for i := range handlers {
		if handlers[i](&wg, ctx, startupTimer, dic) == false {
			cancel()
			break
		}
	}

	// wait for go routines to stop executing
	wg.Wait()
}
