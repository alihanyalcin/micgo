# {project}

### Development Tips
**1.** If you want to add new REST APIs to your microservice, go to **_<project_name>/internal/<service_name>/router.go_** file.

**2.** If you want to update Service, Database, Logging or Startup configuration of your microservice, go to **_<project_name>/cmd/<service_name>/res/configuration.toml_** file.

**3.** If you want to add **new configuration variable**:
* Go to **_<project_name>/internal/pkg/config/types.go_** file, and append your configuration struct:
```go
type TestInfo struct {
	Start string
	End   string
}
```
 * Create **test.go** file under **_<project_name>/internal/pkg/bootstrap/interfaces_** directory, and create an interface for new configuration:
 ```go
type Test interface {
	// GetTestInfo returns a test information.
	GetTestInfo() config.TestInfo
}
```
* Open **_<project_name>/internal/<service_name>/config/config.go_** file, and add new configuration variable to the service, and **GetTestInfo** interface body.
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
 * Go to **_<project_name>/cmd/<service_name>/res/configuration.toml_** file, and append new configuration:
 ```toml
[Test]
Start = "Welcome, service start with new configuration"
End = "Bye, service stop with new configuration"
```
**4.** If you want to create **new handler** with **new configuration variable**:
* Go to **_<project_name>/internal/pkg/bootstrap/handlers_** directory, and create **test** directory.
* Create **test.go** file under **_<project_name>/internal/pkg/bootstrap/handlers/test_** directory.
* Add the lines below to **test.go**:
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
* Then, go to **_<project_name>/cmd/<service_name>/main.go_** file, and append new **BootstrapHandler** to your **microservice**:
```go
test.NewBootstrap(configuration).BootstrapHandler,
```
* Start your service, check log files for new message.

**5.** To add **new database methods**.
* Go to **_<project_name>/internal/pkg/db/interfaces/db.go_** file, and append new method to the interface.
* Go to **_<project_name>/internal/pkg/db/mongo_** directory, and create a **.go** file for **a specific microservice**, then implement your method body.
