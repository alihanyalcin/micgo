package usage

import (
	"fmt"
	"os"
)

var usageStr = `
Usage: %s [options]
Server Options:
    --profile <name>            	Indicate configuration profile other than default
    --confdir                       Specify local configuration directory
Common Options:
    -h, --help                      Show this message
`

// usage will print out the flag options for the server.
func HelpCallback() {
	fmt.Printf(usageStr, os.Args[0])
	os.Exit(0)
}
