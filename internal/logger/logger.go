package logger

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

var logrusLogger *logrus.Logger

func init() {
	logrusLogger = logrus.New()
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})
	logrusLogger.SetLevel(logrus.InfoLevel)
}

func GetLogger() Logger {
	return &logger{
		entry:  logrus.NewEntry(logrusLogger),
		fields: logrus.Fields{},
	}
}

func SetLogLevel(level string) {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		logrusLogger.SetLevel(logrus.DebugLevel)
	case "info":
		logrusLogger.SetLevel(logrus.InfoLevel)
	case "warn":
		logrusLogger.SetLevel(logrus.WarnLevel)
	case "error":
		logrusLogger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrusLogger.SetLevel(logrus.FatalLevel)
	default:
		logrusLogger.SetLevel(logrus.InfoLevel)
	}
}

const (
	fieldError           string = "error"
	fieldErrorMessage    string = "message"
	fieldErrorStackTrace string = "stacktrace"
)

type Logger interface {
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger

	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

type logger struct {
	entry  *logrus.Entry
	fields logrus.Fields
}

func (l *logger) WithField(key string, value interface{}) Logger {
	if key != fieldError {
		l.fields[key] = value
	}
	return l
}

func (l *logger) WithFields(fields map[string]interface{}) Logger {
	for key, value := range fields {
		if key != fieldError {
			l.fields[key] = value
		}
	}
	return l
}

func (l *logger) WithError(err error) Logger {
	l.fields[fieldError] = logrus.Fields{
		fieldErrorMessage:    fmt.Sprintf("%s", err),
		fieldErrorStackTrace: fmt.Sprintf("%+v", err),
	}
	return l
}

func (l *logger) Debug(msg string) {
	l.log(logrus.DebugLevel, msg)
}

func (l *logger) Info(msg string) {
	l.log(logrus.InfoLevel, msg)
}

func (l *logger) Warn(msg string) {
	l.log(logrus.WarnLevel, msg)
}

func (l *logger) Error(msg string) {
	l.log(logrus.ErrorLevel, msg)
}

func (l *logger) Fatal(msg string) {
	l.log(logrus.FatalLevel, msg)
	l.entry.Logger.Exit(1)
}

func (l *logger) log(level logrus.Level, msg string) {
	l.entry.WithFields(logrus.Fields{
		"application": logrus.Fields{
			"name": "gim",
			"gim":  l.fields,
		},
	}).Log(level, msg)
}
