package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(level string) (*Logger, error) {
	// Преобразуем строковый уровень в zapcore.Level
	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, err
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     zap.NewProductionEncoderConfig(),
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
	}

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
	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return err
	}

	l.SugaredLogger.Desugar().Core().Enabled(zapLevel)
	return nil
}
