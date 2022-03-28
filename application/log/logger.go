package log

import (
	"errors"

	"go.uber.org/zap"
)

type ILogger interface {
	Debug(msg string, fields ...zap.Field)
	Debugf(format string, args ...interface{})
	Info(msg string, fields ...zap.Field)
	Infof(format string, args ...interface{})
	Warn(msg string, fields ...zap.Field)
	Warnf(format string, args ...interface{})
	Error(msg string, fields ...zap.Field)
	Errorf(format string, args ...interface{})
	Fatal(msg string, fields ...zap.Field)
	Fatalf(format string, args ...interface{})
}

func init() {
	_, _ = NewLogger()
}

var _logger *Logger

type Logger struct {
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
}

func NewLogger() (*Logger, error) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	sugarLogger := zapLogger.Sugar()
	_logger = &Logger{logger: zapLogger, sugarLogger: sugarLogger}
	return _logger, nil
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.sugarLogger.Debugf(format, args...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.sugarLogger.Fatalf(format, args...)
}

func Debug(msg string, fields ...zap.Field) {
	_logger.Debug(msg, fields...)
}

func Debugf(format string, args ...interface{}) {
	_logger.Debugf(format, args...)
}

func Info(msg string, fields ...zap.Field) {
	_logger.Info(msg, fields...)
}

func Infof(format string, args ...interface{}) {
	_logger.Infof(format, args...)
}

func Warn(msg string, fields ...zap.Field) {
	_logger.Warn(msg, fields...)
}

func Warnf(format string, args ...interface{}) {
	_logger.Warnf(format, args...)
}

func Error(msg string, fields ...zap.Field) {
	_logger.Error(msg, fields...)
}

func Errorf(format string, args ...interface{}) {
	_logger.Errorf(format, args...)
}

func Fatal(msg string, fields ...zap.Field) {
	_logger.Fatal(msg, fields...)
}

func Fatalf(format string, args ...interface{}) {
	_logger.Fatalf(format, args...)
}

func Example() {
	_logger.Info("Info() example")
	_logger.Info("Info() example")
	_logger.Info("Info() with zap.Field example:", zap.String("val", "val1"), zap.Int("val2", 2))
	_logger.Infof("Info() without args")
	_logger.Infof("Info() with args: %s %d %v", "arg1", 23, []int{1, 2, 3})
	_logger.Infof("", "Info() without format and with args. 11", 22, 33.3, map[string]interface{}{"nn": 154})

	err := errors.New("this is an error")
	_logger.Info("", zap.Error(err))
	_logger.Infof("error: %v", err)
	_logger.Infof("", err)

	_logger.Warn("", zap.Error(err))
	_logger.Warnf("error: %v", err)
	_logger.Warnf("", err)

	_logger.Error("", zap.Error(err))
	_logger.Errorf("error: %v", err)
	_logger.Errorf("", err)
}
