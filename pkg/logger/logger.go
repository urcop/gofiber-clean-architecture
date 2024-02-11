package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger
var sugaredLogger *zap.SugaredLogger

func init() {
	namespace := os.Getenv("LOG_NAMESPACE")
	buildMode := os.Getenv("LOG_MODE")

	var config zap.Config
	if buildMode == "production" {
		config = zap.NewProductionConfig()
	} else {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.FunctionKey = "func"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true

	lg, _ := config.Build(zap.AddCallerSkip(1))

	logger = lg.Named(namespace)
	sugaredLogger = lg.Sugar().Named(namespace)
}

func Fatal(msg string, fields ...interface{}) {
	sugaredLogger.Fatalw(msg, fields...)
}

func Error(msg string, fields ...interface{}) {
	sugaredLogger.Errorw(msg, fields...)
}

func Warn(msg string, fields ...interface{}) {
	sugaredLogger.Warnw(msg, fields...)
}

func Info(msg string, fields ...interface{}) {
	sugaredLogger.Infow(msg, fields...)
}

func Debug(msg string, fields ...interface{}) {
	sugaredLogger.Debugw(msg, fields...)
}

/////////////////////////////////////

func FatalZap(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func ErrorZap(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func WarnZap(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func InfoZap(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func DebugZap(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}
