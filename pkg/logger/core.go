package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Core struct {
	core zapcore.Core
}

func NewCore(config *LoggerConfig) *Core {
	encoderConfig := zapcore.EncoderConfig{}

	switch config.EncoderConfigType {
	case "Production":
		encoderConfig = zap.NewProductionEncoderConfig()
	case "Development":
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	default:
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}

	var encoder zapcore.Encoder
	switch config.EncoderType {
	case "JSON":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "Console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	default:
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	ws := config.Writers
	syncers := make([]zapcore.WriteSyncer, len(ws))
	for i := range ws {
		syncers[i] = zapcore.AddSync(ws[i])
	}

	return &Core{core: zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(syncers[0:]...), zapcore.InfoLevel)}
}
