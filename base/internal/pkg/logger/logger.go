/* Package logger provides a client for integration with the support-logging service. The client can also be configured
to write logs to a local file rather than sending them to a service.
*/
package logger

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-kit/kit/log"
)

// LoggingClient defines the interface for logging operations.
type LoggingClient interface {
	// SetLogLevel sets minimum severity log level. If a logging method is called with a lower level of severity than
	// what is set, it will result in no output.
	SetLogLevel(logLevel string) error
	// Debug logs a message at the DEBUG severity level
	Debug(msg string, args ...interface{})
	// Error logs a message at the ERROR severity level
	Error(msg string, args ...interface{})
	// Info logs a message at the INFO severity level
	Info(msg string, args ...interface{})
	// Trace logs a message at the TRACE severity level
	Trace(msg string, args ...interface{})
	// Warn logs a message at the WARN severity level
	Warn(msg string, args ...interface{})
}

type logger struct {
	owningServiceName string
	logTarget         string
	logLevel          *string
	rootLogger        log.Logger
	levelLoggers      map[string]log.Logger
}

type fileWriter struct {
	fileName string
}

// NewClient creates an instance of LoggingClient
func NewClient(owningServiceName string, logTarget string, logLevel string) LoggingClient {
	lc := newClient(owningServiceName, logTarget, logLevel)

	if logTarget == "" {
		lc.Error("logTarget cannot be blank, using stdout only")
	}

	return lc
}

// NewClientStdOut creates an instance of LoggingClient that expects to log to stdout and does not check logTarget
func NewClientStdOut(owningServiceName string, logLevel string) LoggingClient {
	return newClient(owningServiceName, "", logLevel)
}

// newClient is the implementation of the logic required for the factory functions
func newClient(owningServiceName string, logTarget string, logLevel string) logger {
	if !IsValidLogLevel(logLevel) {
		logLevel = InfoLog
	}

	// Set up logging client
	lc := logger{
		owningServiceName: owningServiceName,
		logTarget:         logTarget,
		logLevel:          &logLevel,
	}

	if logTarget != "" { // file based logging
		verifyLogDirectory(lc.logTarget)

		w, err := newFileWriter(lc.logTarget)
		if err != nil {
			stdlog.Fatal(err.Error())
		}
		lc.rootLogger = log.NewLogfmtLogger(io.MultiWriter(os.Stdout, log.NewSyncWriter(w)))
	} else { // HTTP logging OR invalid log target
		lc.rootLogger = log.NewLogfmtLogger(os.Stdout)
	}

	lc.rootLogger = log.WithPrefix(lc.rootLogger, "ts", log.DefaultTimestampUTC,
		"app", owningServiceName, "source", log.Caller(5))

	// Set up the loggers
	lc.levelLoggers = map[string]log.Logger{}

	for _, logLevel := range logLevels() {
		lc.levelLoggers[logLevel] = log.WithPrefix(lc.rootLogger, "level", logLevel)
	}

	return lc
}

// LogLevels returns an array of the possible log levels in order from most to least verbose.
func logLevels() []string {
	return []string{
		TraceLog,
		DebugLog,
		InfoLog,
		WarnLog,
		ErrorLog}
}

func (f *fileWriter) Write(p []byte) (n int, err error) {
	file, err := os.OpenFile(f.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_SYNC, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	_, err = file.WriteString(string(p))
	return len(p), err
}

// IsValidLogLevel checks if `l` is a valid log level
func IsValidLogLevel(l string) bool {
	for _, name := range logLevels() {
		if name == l {
			return true
		}
	}
	return false
}

func newFileWriter(logTarget string) (io.Writer, error) {
	fileWriter := fileWriter{fileName: logTarget}

	return &fileWriter, nil
}

func (lc logger) log(logLevel string, msg string, args ...interface{}) {
	// Check minimum log level
	for _, name := range logLevels() {
		if name == *lc.logLevel {
			break
		}
		if name == logLevel {
			return
		}
	}

	if args == nil {
		args = []interface{}{"msg", msg}
	} else {
		if len(args)%2 == 1 {
			// add an empty string to keep k/v pairs correct
			args = append(args, "")
		}
		if len(msg) > 0 {
			args = append(args, "msg", msg)
		}
	}

	err := lc.levelLoggers[logLevel].Log(args...)
	if err != nil {
		stdlog.Fatal(err.Error())
		return
	}

}

func verifyLogDirectory(path string) {
	prefix, _ := filepath.Split(path)
	//If a path to the log file was specified and it does not exist, create it.
	dir := strings.TrimRight(prefix, "/")
	if len(dir) > 0 {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Println("Creating directory: " + dir)
			os.MkdirAll(dir, 0766)
		}
	}
}

func (lc logger) SetLogLevel(logLevel string) error {
	if IsValidLogLevel(logLevel) == true {
		*lc.logLevel = logLevel

		return nil
	}

	return ErrNotFound{}
}

func (lc logger) Info(msg string, args ...interface{}) {
	lc.log(InfoLog, msg, args...)
}

func (lc logger) Trace(msg string, args ...interface{}) {
	lc.log(TraceLog, msg, args...)
}

func (lc logger) Debug(msg string, args ...interface{}) {
	lc.log(DebugLog, msg, args...)
}

func (lc logger) Warn(msg string, args ...interface{}) {
	lc.log(WarnLog, msg, args...)
}

func (lc logger) Error(msg string, args ...interface{}) {
	lc.log(ErrorLog, msg, args...)
}

// Build the log entry object
func (lc logger) buildLogEntry(logLevel string, msg string, args ...interface{}) LogEntry {
	res := LogEntry{}
	res.Level = logLevel
	res.Message = msg
	res.Args = args
	res.OriginService = lc.owningServiceName

	return res
}
