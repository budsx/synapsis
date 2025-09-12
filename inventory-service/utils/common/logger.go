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
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", filepath.Base(frame.File), frame.Line)
		},
	})
	logger.SetReportCaller(true)

	return &Logger{
		Logger: logger,
	}
}

func (l *Logger) Info(ctx context.Context, message string, args ...any) {
	l.Logger.WithFields(logrus.Fields{
		"[INFO]": args,
	}).Info(message)
}

func (l *Logger) Error(ctx context.Context, message string, args ...any) {
	l.Logger.WithFields(logrus.Fields{
		"[ERROR]": args,
	}).Error(message)
}

func (l *Logger) Debug(ctx context.Context, message string, args ...any) {
	l.Logger.WithFields(logrus.Fields{
		"[DEBUG]": args,
	}).Debug(message)
}

func (l *Logger) Warn(ctx context.Context, message string, args ...any) {
	l.Logger.WithFields(logrus.Fields{
		"[WARNING]": args,
	}).Warn(message)
}
