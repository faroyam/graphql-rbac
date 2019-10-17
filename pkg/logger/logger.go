package logger

import (
	"fmt"
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger for the service
type Logger struct {
	*zap.Logger
}

// NewLogger returns configured zap logger
func NewLogger(level string) (*Logger, error) {
	logLevel, err := strconv.Atoi(level)
	if err != nil {
		logLevel = 1
	}

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.Level(logLevel - 1)),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "ts",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	l, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("initiating logger error: %w", err)
	}

	return &Logger{l}, nil
}
