package pkg

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	log   *logrus.Logger
	level int
}

var Log Logger

const (
	PanicLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func NewLogger() Logger {
	if Log.log == nil {
		Log.log = logrus.New()
	}
	return Log
}

func (l Logger) loggerHandle() (string, int) {
	_, file, line, _ := runtime.Caller(2)
	return file, line
}

func newLine(file string, line int, args ...interface{}) string {
	return fmt.Sprintf("%s:%d %s", file, line, fmt.Sprint(args...))
}

func (l Logger) SetLogger(level int) {
	l.level = level
	l.log.SetLevel(logrus.Level(l.level))
	l.log.SetFormatter(&logrus.TextFormatter{})
	l.log.SetOutput(os.Stdout)

	l.log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func (l Logger) GetLevel() int {
	return int(l.log.Level)
}

func (l Logger) Panic(args ...interface{}) {
	file, line := l.loggerHandle()
	l.log.Panic(newLine(file, line, args...))
}

func (l Logger) Fatal(args ...interface{}) {
	file, line := l.loggerHandle()
	l.log.Fatal(newLine(file, line, args...))
}

func (l Logger) Error(args ...interface{}) {
	file, line := l.loggerHandle()
	l.log.Error(newLine(file, line, args...))
}

func (l Logger) Warn(args ...interface{}) {
	file, line := l.loggerHandle()
	l.log.Warn(newLine(file, line, args...))
}

func (l Logger) Info(args ...interface{}) {
	file, line := l.loggerHandle()
	l.log.Info(newLine(file, line, args...))
}

func (l Logger) Debug(args ...interface{}) {
	file, line := l.loggerHandle()
	l.log.Debug(newLine(file, line, args...))
}

func (l Logger) Trace(args ...interface{}) {
	file, line := l.loggerHandle()
	l.log.Trace(newLine(file, line, args...))
}
