package log

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
	Fatal(args ...interface{})
	Infow(msg string, keysandvalues ...interface{})
	Warnw(msg string, keysandvalues ...interface{})
	Errorw(msg string, keysandvalues ...interface{})
	Panicw(msg string, keysandvalues ...interface{})
	Fatalw(msg string, keysandvalues ...interface{})
}
