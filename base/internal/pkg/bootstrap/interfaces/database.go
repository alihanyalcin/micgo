package interfaces

import "{project}/internal/pkg/config"

// Database interface provides an abstraction for obtaining the database configuration information.
type Database interface {
	// GetDatabaseInfo returns a database information map.
	GetDatabaseInfo() config.DatabaseInfo
}

// CredentialsProvider interface provides an abstraction for obtaining credentials.
type CredentialsProvider interface {
	// GetDatabaseCredentials retrieves database credentials.
	GetDatabaseCredentials(database config.Database) (config.Credentials, error)
}
