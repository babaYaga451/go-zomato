package log

import "go.uber.org/zap"

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() *ZapLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return &ZapLogger{logger: logger.Sugar()}
}

func (z *ZapLogger) Info(args ...interface{}) {
	z.logger.Info(args...)
}

func (z *ZapLogger) Error(args ...interface{}) {
	z.logger.Error(args...)
}

func (z *ZapLogger) Debug(args ...interface{}) {
	z.logger.Debug(args...)
}

func (z *ZapLogger) Warn(args ...interface{}) {
	z.logger.Warn(args...)
}

func (z *ZapLogger) Fatal(args ...interface{}) {
	z.logger.Fatal(args...)
}
func (z *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	z.logger.Infow(msg, keysAndValues...)
}

func (z *ZapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	z.logger.Warnw(msg, keysAndValues...)
}

func (z *ZapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	z.logger.Errorw(msg, keysAndValues...)
}

func (z *ZapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	z.logger.Panicw(msg, keysAndValues...)
}

func (z *ZapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	z.logger.Fatalw(msg, keysAndValues...)
}
