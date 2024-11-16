package logger

import "go.uber.org/zap"

var log *zap.Logger

// InitLogger initializes the global logger instance.
func InitLogger() {
	var err error
	log, err = zap.NewProduction() // Create a production logger
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
}

// Info logs an informational message with optional fields.
func Info(msg string, fields ...zap.Field) {
	if log == nil {
		InitLogger() // Ensure logger is initialized
	}
	log.Info(msg, fields...)
}

// Error logs an error message with optional fields.
func Error(msg string, fields ...zap.Field) {
	if log == nil {
		InitLogger() // Ensure logger is initialized
	}
	log.Error(msg, fields...)
}

// Debug logs a debug message with optional fields.
func Debug(msg string, fields ...zap.Field) {
	if log == nil {
		InitLogger() // Ensure logger is initialized
	}
	log.Debug(msg, fields...)
}

// Warn logs a warning message with optional fields.
func Warn(msg string, fields ...zap.Field) {
	if log == nil {
		InitLogger() // Ensure logger is initialized
	}
	log.Warn(msg, fields...)
}

// Sync flushes any buffered log entries. Should be called on application exit.
func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}
