package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"

	"log"
)

var globalLogger *zap.Logger

func GlobalLogInit() {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder, 
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}
	var encoder zapcore.Encoder
	if os.Getenv("LOG_ENV") == "production" {

		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {

		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(os.Stderr)),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)


	logger := zap.New(core)


	host, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v", err)
	}

	globalLogger = logger.With(zap.String("host", host), zap.Int("pid", os.Getpid()))

	defer globalLogger.Sync()
}
func L() *zap.Logger {
	return globalLogger
}
