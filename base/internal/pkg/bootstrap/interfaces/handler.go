package interfaces

import (
	"context"
	"project/internal/pkg/bootstrap/startup"
	"project/internal/pkg/di"
	"sync"
)

// BootstrapHandler defines the contract each bootstrap handler must fulfill.  Implementation returns true if the
// handler completed successfully, false if it did not.
type BootstrapHandler func(
	wg *sync.WaitGroup,
	ctx context.Context,
	startupTimer startup.Timer,
	dic *di.Container) (success bool)
