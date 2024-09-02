package internals

import (
	"github.com/sirupsen/logrus"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type Log struct {
	logrus *logrus.Logger
}

func NewLogger() *Log {
	l := logrus.New()
	l.SetLevel(logrus.InfoLevel)

	return &Log{
		logrus: l,
	}
}

func (l *Log) Info(args ...interface{}) {
	l.logrus.Info(args...)
}

func (l *Log) Error(args ...interface{}) {
	l.logrus.Error(args...)
}
