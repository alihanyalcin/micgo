package generator

import "os"

var (
	basePath        = os.Getenv("GOPATH") + "/src/github.com/alihanyalcin/micgo/base"
	bin             = "/bin"
	cmd             = "/cmd"
	internal        = "/internal"
	internalService = "/internal/service"
)

var directories = []string{
	bin,
	cmd,
	internal + "/pkg",
	internal + "/pkg/usage",
	internal + "/pkg/telemetry",
	internal + "/pkg/logger",
	internal + "/pkg/di",
	internal + "/pkg/db/interfaces",
	internal + "/pkg/db/mongo",
	internal + "/pkg/db/mongo/models",
	internal + "/pkg/config",
	internal + "/pkg/bootstrap",
	internal + "/pkg/bootstrap/configuration",
	internal + "/pkg/bootstrap/container",
	internal + "/pkg/bootstrap/handlers/database",
	internal + "/pkg/bootstrap/handlers/httpserver",
	internal + "/pkg/bootstrap/handlers/message",
	internal + "/pkg/bootstrap/interfaces",
	internal + "/pkg/bootstrap/logging",
	internal + "/pkg/bootstrap/startup",
}

var files = []string{
	"/VERSION",
	"/version.go",
	"/README.md",
	"/go.mod",
	internal + "/constants.go",
	internal + "/pkg/encoding.go",
	internal + "/pkg/usage/usage.go",
	internal + "/pkg/telemetry/linux_cpu.go",
	internal + "/pkg/telemetry/telemetry.go",
	internal + "/pkg/logger/log_entry.go",
	internal + "/pkg/logger/logger.go",
	internal + "/pkg/di/container.go",
	internal + "/pkg/di/type.go",
	internal + "/pkg/config/types.go",
	internal + "/pkg/db/db.go",
	internal + "/pkg/db/interfaces/db.go",
	internal + "/pkg/db/mongo/client.go",
	internal + "/pkg/bootstrap/bootstrap.go",
	internal + "/pkg/bootstrap/configuration/environment.go",
	internal + "/pkg/bootstrap/configuration/file.go",
	internal + "/pkg/bootstrap/container/configuration.go",
	internal + "/pkg/bootstrap/container/database.go",
	internal + "/pkg/bootstrap/container/logging.go",
	internal + "/pkg/bootstrap/handlers/database/database.go",
	internal + "/pkg/bootstrap/handlers/httpserver/httpserver.go",
	internal + "/pkg/bootstrap/handlers/message/message.go",
	internal + "/pkg/bootstrap/interfaces/configuration.go",
	internal + "/pkg/bootstrap/interfaces/database.go",
	internal + "/pkg/bootstrap/interfaces/handler.go",
	internal + "/pkg/bootstrap/logging/factory.go",
	internal + "/pkg/bootstrap/startup/timer.go",
}
