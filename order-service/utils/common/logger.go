package common

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetReportCaller(false)

	return &Logger{
		Logger: logger,
	}
}

func (l *Logger) Info(ctx context.Context, message string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	filename := filepath.Base(file)
	l.Logger.WithFields(logrus.Fields{
		"args":   args,
		"caller": fmt.Sprintf("%s:%d", filename, line),
	}).Info(message)
}

func (l *Logger) Error(ctx context.Context, message string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	filename := filepath.Base(file)
	l.Logger.WithFields(logrus.Fields{
		"args":   args,
		"caller": fmt.Sprintf("%s:%d", filename, line),
	}).Error(message)
}

func (l *Logger) Debug(ctx context.Context, message string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	filename := filepath.Base(file)
	l.Logger.WithFields(logrus.Fields{
		"args":   args,
		"caller": fmt.Sprintf("%s:%d", filename, line),
	}).Debug(message)
}

func (l *Logger) Warn(ctx context.Context, message string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	filename := filepath.Base(file)
	l.Logger.WithFields(logrus.Fields{
		"args":   args,
		"caller": fmt.Sprintf("%s:%d", filename, line),
	}).Warn(message)
}
