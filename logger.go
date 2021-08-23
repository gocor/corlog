package corlog

import (
	"context"

	"gitlab.com/gocor/corlog/logcfg"
	"go.uber.org/zap"
)

// Logger ...
type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
}

// New ...
func New(ctx context.Context) Logger {
	var l *zap.Logger
	var err error

	config := logcfg.GetSharedConfig()

	if config.IsDebug {
		l, err = zap.NewDevelopment()
	} else {
		l, err = zap.NewProduction()
	}

	if err != nil {
		panic(err)
	}

	filename := getFileName()

	return l.With(
		zap.String("app", config.App),
		zap.String("env", config.Env),
		zap.String("file", filename),
	).Sugar()
}
