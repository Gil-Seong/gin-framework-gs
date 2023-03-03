package logging

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var consoleLogger *zap.Logger
var fileLogger *zap.Logger

func LoggerConfig() {

	logFile := "log/app-%Y-%m-%d-%H.log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(12*time.Hour))
	if err != nil {
		panic(err)
	}
	//파일 로거 정의
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		TimeKey:        "timestamp",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	logCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(rotator),
		zap.InfoLevel)

	fileLogger = zap.New(logCore)

	//콘솔 로거 정의
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	consoleLogger, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	consoleLogger.Info(message, fields...)
	fileLogger.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	consoleLogger.Debug(message, fields...)
	fileLogger.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	consoleLogger.Warn(message, fields...)
	fileLogger.Warn(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	consoleLogger.Panic(message, fields...)
	fileLogger.Panic(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	consoleLogger.Error(message, fields...)
	fileLogger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	consoleLogger.Fatal(message, fields...)
	fileLogger.Fatal(message, fields...)
}
