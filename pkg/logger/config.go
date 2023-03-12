package logger

import "io"

type LoggerConfig struct {
	EncoderType       string
	EncoderConfigType string
	Writers           []io.Writer
}
