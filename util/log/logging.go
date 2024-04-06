package log

import "go.uber.org/zap"

type logger struct {
	name          string
	sugaredLogger *zap.SugaredLogger
}

func (logger logger) Info(args ...any) {
	logger.sugaredLogger.Infoln(args...)
	logger.sugaredLogger.Sync()
}

func (logger logger) Warn(args ...any) {
	logger.sugaredLogger.Warnln(args...)
	logger.sugaredLogger.Sync()
}

func (logger logger) Error(args ...any) {
	logger.sugaredLogger.Errorln(args...)
	logger.sugaredLogger.Sync()
}

func (logger logger) Panic(args ...any) {
	logger.sugaredLogger.Panicln(args...)
}

func New(name string) logger {
	zapLogger, _ := zap.NewProduction()

	return logger{
		name:          name,
		sugaredLogger: zapLogger.Sugar().Named(name),
	}
}
