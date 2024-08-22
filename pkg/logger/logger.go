package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() (*Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	sugar := logger.Sugar()

	return &Logger{sugar}, nil
}

func (l *Logger) SetLevel(level string) error {
	lvl, err := zapcore.ParseLevel(level)
	if err != nil {
		return err
	}
	l.SugaredLogger = l.SugaredLogger.Desugar().WithOptions(zap.IncreaseLevel(lvl)).Sugar()
	return nil
}
