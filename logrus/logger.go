package main

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Log logger
var Log *logrus.Logger

// NewLogger newLogger
func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./info.log",
		logrus.PanicLevel: "./panic.log",
	}

	Log = logrus.New()

	Log.Hooks.Add(lfshook.NewHook(pathMap, &logrus.JSONFormatter{}))

	return Log
}
