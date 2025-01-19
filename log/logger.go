package log

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
)

const (
	logFilePath  = "application.log"
	configPath   = "..\\cofigurationFiles\\config.json"
	defaultLevel = zapcore.InfoLevel // Default log level
)

// Config structure to hold configuration data
type Config struct {
	LogLevel string `json:"log_level"`
}

// Initialize sets up the logger to write logs to a file with the log level from config
func Initialize() error {
	// Load the log level from the configuration file
	logLevel, err := loadLogLevel()
	if err != nil {
		return err
	}

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

	// Set up the encoder (Plain text format with desired fields)
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		MessageKey:   "message",
		CallerKey:    "caller",
		EncodeTime:   millisecondTimeEncoder, // Custom time format
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// Add a log level enabler
	levelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= logLevel // Only enable logs at or above the configured level
	})

	// Create a core with the encoder and file write syncer
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig), // Use plain text encoding
		fileWriteSyncer,                          // Write logs to file
		levelEnabler,                             // Log level from config
	)

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

// Warn logs a warning-level message
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

// loadLogLevel reads the log level from the config.json file
func loadLogLevel() (zapcore.Level, error) {
	file, err := os.Open(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return defaultLevel, nil // Return default log level if config file doesn't exist
		}
		return defaultLevel, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return defaultLevel, err
	}

	// Convert the log level string to zapcore.Level
	return stringToLogLevel(config.LogLevel), nil
}

// stringToLogLevel converts a string to zapcore.Level
func stringToLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return defaultLevel // Default to Info level if invalid input
	}
}
