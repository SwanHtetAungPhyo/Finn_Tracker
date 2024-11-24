package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func GlobalLogInit(){
	config := zap.Config{
		Level: zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: true,
		Encoding: "json",
		EncoderConfig:  zapcore.EncoderConfig{
			TimeKey: "time",
			LevelKey: "level",
			NameKey: "logger",
			CallerKey: "caller",
			MessageKey: "msg",
			StacktraceKey: "stacktrace",
			LineEnding: zapcore.DefaultLineEnding,
			EncodeLevel: zapcore.CapitalLevelEncoder,
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}	

	logger, err := config.Build()
	if err != nil {
		panic("Failed to init the loggger"+ err.Error())
	}
	globalLogger = logger

	defer globalLogger.Sync()
}

func L() *zap.Logger{
	return globalLogger
}