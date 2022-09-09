package kafka

import (
	"go.uber.org/zap"
)

func NewLoggerWrapper(logObj *zap.Logger) *logger {
	return &logger{
		logObj,
	}
}

type logger struct {
	*zap.Logger
}

func (l *logger) Print(v ...interface{}) {
	l.Sugar().Info(v)
}
func (l *logger) Printf(format string, v ...interface{}) {
	l.Sugar().Infof(format, v)
}
func (l *logger) Println(v ...interface{}) {
	l.Sugar().Info(v)
}
