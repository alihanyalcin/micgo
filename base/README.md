# {project}

### Development Tips
**1.** If you want to add new REST APIs to your microservice, go to ```<project_name>/internal/<service_name>/router.go``` file.

**2.** If you want to update Service, Database, Logging or Startup configuration of your microservice, go to ```<project_name>/cmd/<service_name>/res/configuration.toml``` file.

**3.** If you want to add **new configuration variable**:
* Go to ```<project_name>/internal/pkg/config/types.go``` file, and append your configuration struct:
```go
type TestInfo struct {
	Start string
	End   string
}
```
 * Create ```test.go``` file under ```<project_name>/internal/pkg/bootstrap/interfaces``` directory, and create an interface for new configuration:
 ```go
type Test interface {
	// GetTestInfo returns a test information.
	GetTestInfo() config.TestInfo
}
```
* Open ```<project_name>/internal/<service_name>/config/config.go``` file, and add new configuration variable to the service, and **GetTestInfo** interface body.
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
 * Go to ```<project_name>/cmd/<service_name>/res/configuration.toml``` file, and append new configuration:
 ```toml
[Test]
Start = "Welcome, service start with new configuration"
End = "Bye, service stop with new configuration"
```
**4.** If you want to create **new handler** with **new configuration variable**:
* Go to ```<project_name>/internal/pkg/bootstrap/handlers``` directory, and create **test** directory.
* Create ```test.go``` file under ```<project_name>/internal/pkg/bootstrap/handlers/test``` directory.
* Add the lines below to ```test.go```:
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
* Then, go to ```<project_name>/cmd/<service_name>/main.go``` file, and append new **BootstrapHandler** to your **microservice**:
```go
test.NewBootstrap(configuration).BootstrapHandler,
```
* Start your service, check log files for new message.

**5.** To add **new database methods**.
* Go to ```<project_name>/internal/pkg/db/interfaces/db.go``` file, and append new method to the interface.
* Go to ```<project_name>/internal/pkg/db/mongo``` directory, and create a **.go** file for **a specific microservice**, then implement your method body.


**NOTE:** Your project was created with the reorganizing of [this](https://github.com/edgexfoundry/edgex-go) project's source code.