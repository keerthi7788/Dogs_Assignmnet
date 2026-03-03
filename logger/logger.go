package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Global Logger variable
var Logger *zap.Logger

// InitLogger initializes the logger and configures its settings
func InitLogger() {
	// Custom encoder config for tab-separated output
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",                                      // Use tab as a line ending
		EncodeLevel:    zapcore.CapitalLevelEncoder,               // Uppercase log levels
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.RFC3339), // Custom time format
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // Short caller
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Using stdout for output
	writer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)
	Logger = zap.New(core, zap.AddCaller(), zap.Development(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Info(msg string, args ...interface{}) {
	Logger.Info(msg, ConvertArgsToFields(args...)...)
}

func Error(msg string, args ...interface{}) {
	Logger.Error(msg, ConvertArgsToFields(args...)...)
}

func Debug(msg string, args ...interface{}) {
	Logger.Debug(msg, ConvertArgsToFields(args...)...)
}

func Fatal(msg string, args ...interface{}) {
	Logger.Fatal(msg, ConvertArgsToFields(args...)...)
}

func Warn(msg string, args ...interface{}) {
	Logger.Warn(msg, ConvertArgsToFields(args...)...)
}

func Panic(msg string, args ...interface{}) {
	Logger.Panic(msg, ConvertArgsToFields(args...)...)
}

func ConvertArgsToFields(args ...interface{}) []zap.Field {
	fields := make([]zap.Field, len(args))
	for i, arg := range args {
		fields[i] = convertToField(arg)
	}
	return fields
}

// convertToField converts an argument to a zap.Field based on its type
func convertToField(arg interface{}) zap.Field {
	switch v := arg.(type) {
	case string:
		return zap.String("string", v)
	case int:
		return zap.Int("int", v)
	case int64:
		return zap.Int64("int64", v)
	case float64:
		return zap.Float64("float64", v)
	case bool:
		return zap.Bool("bool", v)
	case error:
		return zap.Error(v)
	case rune:
		return zap.String("rune", string(v))
	default:
		return zap.Any("any", v)
	}
}
