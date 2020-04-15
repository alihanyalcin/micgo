# {project}

### Development Tips
1. If you want to add new API, go to <project_name>/internal/<service_name>/router.go
2. If you want to update service configuration, go to <project_name>/cmd/<service_name>/res/configuration.toml
3. If you want to add new configuration variable:
* Go to <project_name>/internal/pkg/config/types.go, and append your configuration struct:
```go
type TestInfo struct {
	Start string
	End   string
}
```
 * Go to <project_name>/cmd/<service_name>/res/configuration.toml, and append new configuration:
 ```toml
[Test]
Start = "Welcome, service start with new configuration"
End = "Bye, service stop with new configuration"
```
 * Create test.go file under <project_name>/internal/pkg/bootstrap/interfaces/ directory, and append lines:
 ```go
type Test interface {
	// GetTestInfo returns a test information.
	GetTestInfo() config.TestInfo
}
```
* Open <project_name>/internal/<service_name>/config/config.go, and add new configuration variable and append GetTestInfo interface.
```go
// Add Test variable
type ConfigurationStruct struct {
	Service   config.ServiceInfo
	Logging   config.LoggingInfo
	Startup   config.StartupInfo
	Databases config.DatabaseInfo
	Test      config.TestInfo
}

// Append GetTestInfo
// GetTestInfo returns a test information.
func (c *ConfigurationStruct) GetTestInfo() config.TestInfo {
	return c.Test
}
```
4. Create new handler with new configuration variable.
* Go to <project_name>/internal/pkg/bootstrap/handlers, and create test directory and test.go.
* Append codes:
```go
type TestMessage struct {
	Message interfaces.Test
}

func NewBootstrap(test interfaces.Test) TestMessage {
	return TestMessage{
		Message: test,
	}
}

func (s TestMessage) BootstrapHandler(
	wg *sync.WaitGroup,
	ctx context.Context,
	startupTimer startup.Timer,
	dic *di.Container) bool {

	loggingClient := container.LoggingClientFrom(dic.Get)
	loggingClient.Info("Test Configuration Start:"+s.Message.GetTestInfo().Start)

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-ctx.Done()
		loggingClient.Info("Test Configuration End:"+s.Message.GetTestInfo().End)
	}()

	return true
}
```
* then, go to <project_name>/cmd/<service_name>/main.go. And append new BootstrapHandler:
```go
test.NewBootstrap(configuration).BootstrapHandler,
```
* start your service, check log files for new message

5. To add new database methods.
* Go to <project_name>/internal/pkg/db/interfaces/db.go, and append new method.
* Go to <project_name>/internal/pkg/db/mongo/, and create a go file for specific database, then implement your method body.