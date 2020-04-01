package httpserver

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"project/internal/pkg/bootstrap/container"
	"project/internal/pkg/bootstrap/startup"
	"project/internal/pkg/di"
	"strconv"
	"sync"
	"time"
)

// HttpServer contains references to dependencies required by the http server implementation.
type HttpServer struct {
	router    *mux.Router
	isRunning bool
}

// NewBootstrap is a factory method that returns an initialized HttpServer receiver struct.
func NewBootstrap(router *mux.Router) HttpServer {
	return HttpServer{
		router:    router,
		isRunning: false,
	}
}

// IsRunning returns whether or not the http server is running.  It is provided to support delayed shutdown of
// any resources required to successfully process http requests until after all outstanding requests have been
// processed (e.g. a database connection).
func (h *HttpServer) IsRunning() bool {
	return h.isRunning
}

// BootstrapHandler fulfills the BootstrapHandler contract.  It creates two go routines -- one that executes ListenAndServe()
// and another that waits on closure of a context's done channel before calling Shutdown() to cleanly shut down the
// http server.
func (h *HttpServer) BootstrapHandler(
	wg *sync.WaitGroup,
	ctx context.Context,
	startupTimer startup.Timer,
	dic *di.Container) bool {

	bootstrapConfig := container.ConfigurationFrom(dic.Get).GetBootstrap()
	addr := bootstrapConfig.Service.Host + ":" + strconv.Itoa(bootstrapConfig.Service.Port)
	timeout := time.Millisecond * time.Duration(bootstrapConfig.Service.Timeout)

	server := &http.Server{
		Addr:         addr,
		Handler:      h.router,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	loggingClient := container.LoggingClientFrom(dic.Get)
	loggingClient.Info("Web server starting (" + addr + ")")

	wg.Add(1)
	go func() {
		defer wg.Done()

		h.isRunning = true
		server.ListenAndServe()
		loggingClient.Info("Web server stopped")
		h.isRunning = false
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		loggingClient.Info("Web Server shutting down")
		server.Shutdown(context.Background())
		loggingClient.Info("Web Server shut down")
	}()

	return true
}
