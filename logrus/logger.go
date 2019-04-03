package main

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
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

// NewLoggerWithRotate NewLoggerWithRotate
func NewLoggerWithRotate() *logrus.Logger {
	if Log != nil {
		return Log
	}

	path := "./test.log"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(path),               // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(60*time.Second),       // 文件最大保存时间
		rotatelogs.WithRotationTime(20*time.Second), // 日志切割时间间隔
	)

	pathMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.PanicLevel: writer,
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	return Log
}
