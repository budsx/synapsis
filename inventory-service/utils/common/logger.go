package common

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	logger.SetReportCaller(false)

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.SetLevel(logrus.InfoLevel)

	return &Logger{
		logger: logger,
	}
}

func (l *Logger) GetLogger() *logrus.Entry {
	if _, file, line, ok := runtime.Caller(2); ok {
		filename := filepath.Base(file)
		return l.logger.WithField("file", fmt.Sprintf("%s:%d", filename, line))
	}
	return l.logger.WithField("file", "unknown")
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.GetLogger().Infof(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.GetLogger().Errorf(format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	l.GetLogger().Fatalf(format, v...)
}

func (l *Logger) Debug(format string, v ...interface{}) {
	l.GetLogger().Debugf(format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.GetLogger().Warnf(format, v...)
}

func (l *Logger) WithField(key string, value interface{}) *logrus.Entry {
	return l.logger.WithField(key, value)
}

func (l *Logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.logger.WithFields(fields)
}

func (l *Logger) SetLevel(level logrus.Level) {
	l.logger.SetLevel(level)
}
