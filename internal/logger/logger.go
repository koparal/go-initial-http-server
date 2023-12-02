package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"initial-http-server/internal/config"
)

var Logger *zap.Logger

func InitLogger(loggingConfig config.LoggingConfig) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = zap.NewAtomicLevelAt(getZapLevel(loggingConfig.Level))
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.OutputPaths = []string{loggingConfig.Filename}

	var err error
	Logger, err = zapConfig.Build()
	if err != nil {
		panic(err)
	}
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
