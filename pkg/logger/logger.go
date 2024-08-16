// File: pkg/logger/logger.go

package logger

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	log  zerolog.Logger
	once sync.Once
)

// Init initializes the logger
func Init(level string) error {
	var err error
	once.Do(func() {
		// Parse log level
		lvl, err := zerolog.ParseLevel(level)
		if err != nil {
			return
		}

		// Set global log level
		zerolog.SetGlobalLevel(lvl)

		// Create logger
		log = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	})
	return err
}

// Info logs a message at InfoLevel
func Info(message string, fields ...interface{}) {
	log.Info().Fields(fields).Msg(message)
}

// Error logs a message at ErrorLevel
func Error(message string, fields ...interface{}) {
	log.Error().Fields(fields).Msg(message)
}

// Debug logs a message at DebugLevel
func Debug(message string, fields ...interface{}) {
	log.Debug().Fields(fields).Msg(message)
}

// Warn logs a message at WarnLevel
func Warn(message string, fields ...interface{}) {
	log.Warn().Fields(fields).Msg(message)
}

// Fatal logs a message at FatalLevel and then calls os.Exit(1)
func Fatal(message string, fields ...interface{}) {
	log.Fatal().Fields(fields).Msg(message)
}

// With creates a child logger and adds structured context to it
func With(fields ...interface{}) zerolog.Logger {
	return log.With().Fields(fields).Logger()
}

// Sync is a no-op function to maintain API compatibility
func Sync() error {
	return nil
}
