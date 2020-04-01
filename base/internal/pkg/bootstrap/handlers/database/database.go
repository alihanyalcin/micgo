package database

import (
	"context"
	"fmt"
	"project/internal/pkg/bootstrap/container"
	"project/internal/pkg/bootstrap/interfaces"
	"project/internal/pkg/bootstrap/startup"
	"project/internal/pkg/db"
	dbInterfaces "project/internal/pkg/db/interfaces"
	"project/internal/pkg/db/mongo"
	"project/internal/pkg/di"
	"project/internal/pkg/logger"
	"sync"
	"time"
)

// httpServer defines the contract used to determine whether or not the http httpServer is running.
type httpServer interface {
	IsRunning() bool
}

// Database contains references to dependencies required by the database bootstrap implementation.
type Database struct {
	httpServer httpServer
	database   interfaces.Database
}

// NewDatabase is a factory method that returns an initialized Database receiver struct.
func NewDatabase(httpServer httpServer, database interfaces.Database) Database {
	return Database{
		httpServer: httpServer,
		database:   database,
	}
}

// Return the dbClient interface
func (d Database) newDBClient(loggingClient logger.LoggingClient) (dbInterfaces.DBClient, error) {
	databaseInfo := d.database.GetDatabaseInfo()["Primary"]
	switch databaseInfo.Type {
	case db.MongoDB:
		return mongo.NewClient(
			db.Configuration{
				Host:         databaseInfo.Host,
				Port:         databaseInfo.Port,
				Timeout:      databaseInfo.Timeout,
				DatabaseName: databaseInfo.Name,
				Username:     databaseInfo.Username,
				Password:     databaseInfo.Password,
			})
	default:
		return nil, db.ErrUnsupportedDatabase
	}
}

// BootstrapHandler fulfills the BootstrapHandler contract and initializes the database.
func (d Database) BootstrapHandler(
	wg *sync.WaitGroup,
	context context.Context,
	startupTimer startup.Timer,
	dic *di.Container) bool {

	loggingClient := container.LoggingClientFrom(dic.Get)

	// initialize database
	var dbClient dbInterfaces.DBClient
	for startupTimer.HasNotElapsed() {
		var err error
		dbClient, err = d.newDBClient(loggingClient)
		if err == nil {
			break
		}
		dbClient = nil
		loggingClient.Warn(fmt.Sprintf("couldn't create database client: %v", err.Error()))
		startupTimer.SleepForInterval()
	}

	if dbClient == nil {
		return false
	}

	dic.Update(di.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get di.Get) interface{} {
			return dbClient
		},
	})
	loggingClient.Info("Database connected.")

	wg.Add(1)
	go func() {
		defer wg.Done()

		<-context.Done()
		for {
			// wait for httpServer to stop running (e.g. handling requests) before closing the database connection.
			if d.httpServer.IsRunning() == false {
				dbClient.CloseSession()
				break
			}
			time.Sleep(time.Second)
		}
		loggingClient.Info("Database disconnected")
	}()

	return true
}
