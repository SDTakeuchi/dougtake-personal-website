package log

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

type (
	logger struct {
		loggerImpl *slog.Logger
	}
	Logger interface {
		Debug(msg string)
		Debugf(format string, args ...interface{})
		Info(msg string)
		Infof(format string, args ...interface{})
		Warn(msg string)
		Warnf(format string, args ...interface{})
		Error(msg string)
		Errorf(format string, args ...interface{})
		Fatal(msg string)
		Fatalf(format string, args ...interface{})
	}
)

func NewLogger() Logger {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove time.
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	l := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource:   true,
				ReplaceAttr: replace,
			},
		),
	)
	return &logger{loggerImpl: l}
}

func (l *logger) Debug(msg string) {
	l.loggerImpl.Debug(msg)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.loggerImpl.Debug(format, args...)
}

func (l *logger) Info(msg string) {
	l.loggerImpl.Info(msg)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.loggerImpl.Info(format, args...)
}

func (l *logger) Warn(msg string) {
	l.loggerImpl.Warn(msg)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.loggerImpl.Warn(format, args...)
}

func (l *logger) Error(msg string) {
	l.loggerImpl.Error(msg)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.loggerImpl.Error(format, args...)
}

func (l *logger) Fatal(msg string) {
	l.loggerImpl.Error(msg)
	fmt.Println("fatal error occurred")
	os.Exit(1)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}
