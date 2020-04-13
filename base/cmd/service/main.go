package main

import (
	"flag"
	"{project}"
	"{project}/internal"
	"{project}/internal/pkg/bootstrap"
	"{project}/internal/pkg/bootstrap/handlers/database"
	"{project}/internal/pkg/bootstrap/handlers/httpserver"
	"{project}/internal/pkg/bootstrap/handlers/message"
	"{project}/internal/pkg/bootstrap/interfaces"
	"{project}/internal/pkg/bootstrap/startup"
	"{project}/internal/pkg/di"
	"{project}/internal/pkg/usage"
	"{project}/internal/{servicename}"
	"{project}/internal/{servicename}/config"
)

func main() {
	startupTimer := startup.NewStartUpTimer(internal.BootRetrySecondsDefault, internal.BootTimeoutSecondsDefault)

	var configDir, profileDir string
	flag.StringVar(&profileDir, "profile", "", "Specify a profile other than default.")
	flag.StringVar(&configDir, "confdir", "", "Specify local configuration directory.")

	flag.Usage = usage.HelpCallback
	flag.Parse()

	configuration := &config.ConfigurationStruct{}
	dic := di.NewContainer(di.ServiceConstructorMap{})

	httpServer := httpserver.NewBootstrap({servicename}.LoadRestRoutes())
	bootstrap.Run(
		configDir,
		profileDir,
		internal.ConfigFileName,
		"{servicename}",
		configuration,
		startupTimer,
		dic,
		[]interfaces.BootstrapHandler{
			database.NewDatabase(&httpServer, configuration).BootstrapHandler,
			{servicename}.BootstrapHandler,
			httpServer.BootstrapHandler,
			message.NewBootstrap("{servicename}", {project}.Version).BootstrapHandler,
		})
}
