package message

import (
	"context"
	"fmt"
	"sync"

	"project/internal/pkg/bootstrap/container"
	"project/internal/pkg/bootstrap/startup"
	"project/internal/pkg/di"
)

// StartMessage contains references to dependencies required by the start message handler.
type StartMessage struct {
	serviceKey string
	version    string
}

// NewBootstrap is a factory method that returns an initialized StartMessage receiver struct.
func NewBootstrap(serviceKey, version string) StartMessage {
	return StartMessage{
		serviceKey: serviceKey,
		version:    version,
	}
}

// BootstrapHandler fulfills the BootstrapHandler contract.  It creates no go routines.  It logs a "standard" set of
// messages when the service first starts up successfully.
func (s StartMessage) BootstrapHandler(
	wg *sync.WaitGroup,
	ctx context.Context,
	startupTimer startup.Timer,
	dic *di.Container) bool {

	loggingClient := container.LoggingClientFrom(dic.Get)
	loggingClient.Info("Service dependencies resolved...")
	loggingClient.Info(fmt.Sprintf("Starting %s %s ", s.serviceKey, s.version))

	bootstrapConfig := container.ConfigurationFrom(dic.Get).GetBootstrap()
	if len(bootstrapConfig.Service.StartupMsg) > 0 {
		loggingClient.Info(bootstrapConfig.Service.StartupMsg)
	}

	loggingClient.Info("Service started in: " + startupTimer.SinceAsString())

	return true
}
