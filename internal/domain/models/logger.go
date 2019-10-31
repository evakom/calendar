/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 30.10.2019 15:18
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package models implements base models.
package models

import (
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

// Logger is the base struct for all loggers.
type Logger struct {
	logger *log.Logger
}

var lg Logger

// GetLogger returns global logger.
func (l Logger) GetLogger() Logger {
	return lg
}

// NewLogger inits logger instance.
func NewLogger(level string, output io.Writer) {
	lg = Logger{logger: log.New()}
	if level == "none" {
		lg.logger.SetOutput(ioutil.Discard)
	} else {
		lg.logger.SetOutput(output)
	}
	switch level {
	case "debug":
		lg.logger.SetLevel(log.DebugLevel)
	case "info":
		lg.logger.SetLevel(log.InfoLevel)
	case "warn":
		lg.logger.SetLevel(log.WarnLevel)
	default:
		lg.logger.SetLevel(log.ErrorLevel)
	}
}

// Debug writes debug level to output.
func (l Logger) Debug(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Info writes info level to output.
func (l Logger) Info(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warn writes warn level to output.
func (l Logger) Warn(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Error writes error level to output.
func (l Logger) Error(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// WithRequest logs all request fields.
//func (l Logger) WithRequest(req *http.Request) Logger {
//  return l.logger.WithField("addr", "qqqqqqq")
//	return l.logger.WithFields(RequestFields(req))
//}
