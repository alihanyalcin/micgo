package generator

import "os"

var (
	basePath = os.Getenv("GOPATH") + "/src/github.com/alihanyalcin/micgo/base/"
	files    = map[string]string{
		"README.md":             basePath + "README.md",
		"go.mod":                basePath + "go.mod",
		"VERSION":               basePath + "VERSION",
		"version.go":            basePath + "version.go",
		"internal/constants.go": basePath + "internal/constants.go",
	}
)
