// pkg/logger/logger.go

package logger

import (
    "go.uber.org/zap"
)

var log *zap.Logger

func Init(level string) {
    var err error

    config := zap.NewProductionConfig()

    err = config.Level.UnmarshalText([]byte(level))
    if err != nil {
        panic(err)
    }

    log, err = config.Build(zap.AddCallerSkip(1))
    if err != nil {
        panic(err)
    }
}

func Info(message string, fields ...zap.Field) {
    log.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
    log.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
    log.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
    log.Fatal(message, fields...)
}