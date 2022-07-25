package echologrus

import (
	"io"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// MyLogger extend logrus.Logger
type MyLogger struct {
	*logrus.Logger
}

// Singleton logger
var singletonLogger = &MyLogger{
	Logger: logrus.New(),
}

// Logger return singleton logger
func Logger() *MyLogger {
	return singletonLogger
}

func Print(i ...interface{}) {
	singletonLogger.Print(i...)
}

func Printf(format string, i ...interface{}) {
	singletonLogger.Printf(format, i...)
}

// Not used
func Printj(j log.JSON) {
	singletonLogger.Printj(j)
}

func Debug(i ...interface{}) {
	singletonLogger.Debug(i...)
}

// Not used
func Debugf(format string, args ...interface{}) {
	singletonLogger.Debugf(format, args...)
}

// Not used
func Debugj(j log.JSON) {
	singletonLogger.Debugj(j)
}

func Info(i ...interface{}) {
	singletonLogger.Info(i...)
}

// Not used
func Infof(format string, args ...interface{}) {
	singletonLogger.Infof(format, args...)
}

// Not used
func Infoj(j log.JSON) {
	singletonLogger.Infoj(j)
}

func Warn(i ...interface{}) {
	singletonLogger.Warn(i...)
}

// Not used
func Warnf(format string, args ...interface{}) {
	singletonLogger.Warnf(format, args...)
}

// Not used
func Warnj(j log.JSON) {
	singletonLogger.Warnj(j)
}

func Error(i ...interface{}) {
	singletonLogger.Error(i...)
}

// Not used
func Errorf(format string, args ...interface{}) {
	singletonLogger.Errorf(format, args...)
}

// Not used
func Errorj(j log.JSON) {
	singletonLogger.Errorj(j)
}

func Fatal(i ...interface{}) {
	singletonLogger.Fatal(i...)
}

// Not used
func Fatalf(format string, args ...interface{}) {
	singletonLogger.Fatalf(format, args...)
}

// Not used
func Fatalj(j log.JSON) {
	singletonLogger.Fatalj(j)
}

func Panic(i ...interface{}) {
	singletonLogger.Panic(i...)
}

// Not used
func Panicf(format string, args ...interface{}) {
	singletonLogger.Panicf(format, args...)
}

// Not used
func Panicj(j log.JSON) {
	singletonLogger.Panicj(j)
}

// To logrus.Level
func toLogrusLevel(level log.Lvl) logrus.Level {
	switch level {
	case log.DEBUG:
		return logrus.DebugLevel
	case log.INFO:
		return logrus.InfoLevel
	case log.WARN:
		return logrus.WarnLevel
	case log.ERROR:
		return logrus.ErrorLevel
	}

	return logrus.InfoLevel
}

// To Echo.log.lvl
func ToEchoLevel(level logrus.Level) log.Lvl {
	switch level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	}

	return log.OFF
}

// Output return logger io.Writer
func (l *MyLogger) Output() io.Writer {
	return l.Out
}

// SetOutput logger io.Writer
func (l *MyLogger) SetOutput(w io.Writer) {
	l.Out = w
}

// Level return logger level
func (l *MyLogger) Level() log.Lvl {
	return ToEchoLevel(l.Logger.Level)
}

// SetLevel logger level
func (l *MyLogger) SetLevel(v log.Lvl) {
	l.Logger.Level = toLogrusLevel(v)
}

// SetHeader logger header
// Managed by Logrus itself
// This function do nothing
func (l *MyLogger) SetHeader(h string) {
	// do nothing
}

// Formatter return logger formatter
func (l *MyLogger) Formatter() logrus.Formatter {
	return l.Logger.Formatter
}

// SetFormatter logger formatter
// Only support logrus formatter
func (l *MyLogger) SetFormatter(formatter logrus.Formatter) {
	l.Logger.Formatter = formatter
}

// Prefix return logger prefix
// This function do nothing
func (l *MyLogger) Prefix() string {
	return ""
}

// SetPrefix logger prefix
// This function do nothing
func (l *MyLogger) SetPrefix(p string) {
	// Do nothing
}

func (l *MyLogger) Print(i ...interface{}) {
	l.Logger.Print(i...)
}

// Used in gormlogrus
func (l *MyLogger) Printf(format string, args ...interface{}) {
	l.Logger.Printf(format, args...)
}

// Not used
func (l *MyLogger) Printj(j log.JSON) {
	// Do nothing
}

func (l *MyLogger) Debug(i ...interface{}) {
	l.Logger.Debug(i...)
}

// Not used
func (l *MyLogger) Debugf(format string, args ...interface{}) {
	// Do nothing
}

// Not used
func (l *MyLogger) Debugj(j log.JSON) {
	// Do nothing
}

func (l *MyLogger) Info(i ...interface{}) {
	l.Logger.Info(i...)
}

// Not used
func (l *MyLogger) Infof(format string, args ...interface{}) {
	// Do nothing
}

// Not used
func (l *MyLogger) Infoj(j log.JSON) {
	// Do nothing
}

func (l *MyLogger) Warn(i ...interface{}) {
	l.Logger.Warn(i...)
}

// Not used
func (l *MyLogger) Warnf(format string, args ...interface{}) {
	// Do nothing
}

// Not used
func (l *MyLogger) Warnj(j log.JSON) {
	// Do nothing
}

func (l *MyLogger) Error(i ...interface{}) {
	l.Logger.Error(i...)
}

// Not used
func (l *MyLogger) Errorf(format string, args ...interface{}) {
	// Do nothing
}

// Not used
func (l *MyLogger) Errorj(j log.JSON) {
	// Do nothing
}

func (l *MyLogger) Fatal(i ...interface{}) {
	l.Logger.Fatal(i...)
}

// Not used
func (l *MyLogger) Fatalf(format string, args ...interface{}) {
	// Do nothing
}

// Not used
func (l *MyLogger) Fatalj(j log.JSON) {
	// Do nothing
}

func (l *MyLogger) Panic(i ...interface{}) {
	l.Logger.Panic(i...)
}

// Not used
func (l *MyLogger) Panicf(format string, args ...interface{}) {
	// Do nothing
}

// Not used
func (l *MyLogger) Panicj(j log.JSON) {
	// Do nothing
}

// Log output settings for golang-migrate
func (l *MyLogger) Verbose() bool {
	return true
}
