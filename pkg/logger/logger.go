package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var Logger zerolog.Logger

func InitLogger(level, output string) error {
	var logOutput zerolog.ConsoleWriter

	if output == "" {
		logOutput = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
	} else {
		file, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		logOutput = zerolog.ConsoleWriter{Out: file, TimeFormat: "2006-01-02 15:04:05", NoColor: true}
	}
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}
	Logger = zerolog.New(logOutput).With().Timestamp().Logger().Level(logLevel)
	return nil
}
func Debug(msg string, fields map[string]interface{}) {
	Logger.Debug().Fields(fields).Msg(msg)
}

func Info(msg string, fields map[string]interface{}) {
	Logger.Info().Fields(fields).Msg(msg)
}

func Warn(msg string, fields map[string]interface{}) {
	Logger.Warn().Fields(fields).Msg(msg)
}

func Error(msg string, fields map[string]interface{}) {
	Logger.Error().Fields(fields).Msg(msg)
}
func Fatal(msg string, fields map[string]interface{}) {
	Logger.Error().Fields(fields).Msg(msg)
	os.Exit(1)
}
