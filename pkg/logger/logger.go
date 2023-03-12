package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field
type Level = zapcore.Level

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

type CatLogger struct {
	logger *zap.Logger
}

func NewLogger(core *Core) *CatLogger {
	return &CatLogger{
		logger: zap.New(core.core),
	}
}

func (log *CatLogger) Log(lvl Level, msg string, fields ...Field) {
	log.logger.Log(lvl, msg, fields...)
}

func (log *CatLogger) Debug(msg string, fields ...Field) {
	log.logger.Debug(msg, fields...)
}

func (log *CatLogger) Info(msg string, fields ...Field) {
	log.logger.Info(msg, fields...)
}
func (log *CatLogger) Warn(msg string, fields ...Field) {
	log.logger.Warn(msg, fields...)
}

func (log *CatLogger) Error(msg string, fields ...Field) {
	log.logger.Error(msg, fields...)
}

func (log *CatLogger) DPanic(msg string, fields ...Field) {
	log.logger.DPanic(msg, fields...)
}

func (log *CatLogger) Panic(msg string, fields ...Field) {
	log.logger.Panic(msg, fields...)
}

func (log *CatLogger) Fatal(msg string, fields ...Field) {
	log.logger.Fatal(msg, fields...)
}
