package log

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
)

type (
	loggerImpl struct {
		logger *slog.Logger
	}
	Logger interface {
		// Debug logs a message at the debug level.
		Debug(msg string)

		// Debugf logs a formatted message at the debug level.
		Debugf(format string, args ...interface{})

		// Info logs a message at the info level.
		Info(msg string)

		// Infof logs a formatted message at the info level.
		Infof(format string, args ...interface{})

		// Warn logs a message at the warn level.
		Warn(msg string)

		// Warnf logs a formatted message at the warn level.
		Warnf(format string, args ...interface{})

		// Error logs a message at the error level.
		Error(msg string)

		// Errorf logs a formatted message at the error level.
		Errorf(format string, args ...interface{})

		// Fatal logs a message at the error level and exits the program.
		Fatal(msg string)

		// Fatalf logs a formatted message at the error level and exits the program.
		Fatalf(format string, args ...interface{})
	}
)

var defaultLogLevel = new(slog.LevelVar) // Info by default

func NewLogger() Logger {
	l := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: false,
				Level:     defaultLogLevel,
			},
		),
	)
	return &loggerImpl{logger: l}
}

// showStackTrace returns a string representation of the current stack trace.
// The depth parameter controls how many stack frames to include in the output.
// A depth of 0 includes all stack frames.
func showStackTrace(depth int) string {
	// first three stacktrace show the logger's stacktrace, so skip them
	defaultSkipCount := 3
	depth = depth + defaultSkipCount
	i := defaultSkipCount
	stacktrace := ""
	for {
		pt, file, line, ok := runtime.Caller(i)
		if !ok || (depth != defaultSkipCount && i >= depth) {
			break
		}
		funcName := runtime.FuncForPC(pt).Name()
		stacktrace += fmt.Sprintf("file=%s, line=%d, func=%v\n", file, line, funcName)
		i += 1
	}
	return stacktrace
}

func (l *loggerImpl) Debug(msg string) {
	l.logger.Debug(msg, "stacktrace", showStackTrace(1))
}

func (l *loggerImpl) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *loggerImpl) Info(msg string) {
	l.logger.Info(msg, "stacktrace", showStackTrace(1))
}

func (l *loggerImpl) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *loggerImpl) Warn(msg string) {
	l.logger.Warn(msg, "stacktrace", showStackTrace(1))
}

func (l *loggerImpl) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *loggerImpl) Error(msg string) {
	l.logger.Error(msg, "stacktrace", showStackTrace(0))
}

func (l *loggerImpl) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *loggerImpl) Fatal(msg string) {
	l.Error(msg)
	os.Exit(1)
}

func (l *loggerImpl) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}
