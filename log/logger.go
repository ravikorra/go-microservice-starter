package log

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
)

const logFilePath string = "application.log"

// Initialize sets up the logger to write logs to a file
func Initialize() error {
	// Create log directory if it doesn't exist
	logDir := filepath.Dir(logFilePath)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if mkdirErr := os.MkdirAll(logDir, 0755); mkdirErr != nil {
			return mkdirErr
		}
	}

	// Create a file write syncer
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	fileWriteSyncer := zapcore.AddSync(file)

	// Set up the encoder (JSON format)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "timestamp",
		LevelKey:   "level",
		MessageKey: "message",
		CallerKey:  "caller",
		EncodeTime: millisecondTimeEncoder, // Custom time format
		//EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // Use plain text encoding
		fileWriteSyncer,                          // Write logs to file
		zapcore.DebugLevel,                       // Log level
	)

	// Set log level
	// core := zapcore.NewCore(
	// 	zapcore.NewJSONEncoder(encoderConfig),
	// 	fileWriteSyncer,
	// 	zapcore.DebugLevel, // Change this to zapcore.InfoLevel, zapcore.ErrorLevel, etc., as needed
	// )

	// Build the logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return nil
}

// Debug logs a debug-level message
func Debug(message string, fields ...zap.Field) {
	Logger.Debug(message, fields...)
}

// Info logs an info-level message
func Info(message string, fields ...zap.Field) {
	Logger.Info(message, fields...)
}

// Error logs an error-level message
func Error(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
}

// Warn logs an warning-level message
func Warn(message string, fields ...zap.Field) {
	Logger.Warn(message, fields...)
}

// Custom time encoder for human-readable format with milliseconds
func millisecondTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01- 19 15:04:05.000"))
}

// Sync flushes any buffered log entries
func Sync() {
	_ = Logger.Sync()
}
