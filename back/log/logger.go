package log

import (
	"log"
	"os"
	"sync"
)

// Logger is the interface for logging operations.
type Logger interface {
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
	Fatal(format string, v ...any)
	Panic(format string, v ...any)
}

type defaultLogger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
	panicLogger *log.Logger
}

var (
	once   sync.Once
	logger *defaultLogger
)

// NewLogger initializes a new logger instance if it hasn't been initialized yet.
// For simplicity, it currently logs to os.Stdout for all levels.
// In a real application, you might configure different outputs (e.g., files, syslog)
// and different levels based on environment variables or configuration files.
func NewLogger() Logger {
	once.Do(func() {
		logger = &defaultLogger{
			debugLogger: log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime),
			infoLogger:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
			warnLogger:  log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime),
			errorLogger: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime),
			fatalLogger: log.New(os.Stderr, "[FATAL] ", log.Ldate|log.Ltime),
			panicLogger: log.New(os.Stderr, "[PANIC] ", log.Ldate|log.Ltime),
		}
	})
	return logger
}

func (l *defaultLogger) Debug(format string, v ...any) {
	l.debugLogger.Printf(format, v...)
}

func (l *defaultLogger) Info(format string, v ...any) {
	l.infoLogger.Printf(format, v...)
}

func (l *defaultLogger) Warn(format string, v ...any) {
	l.warnLogger.Printf(format, v...)
}

func (l *defaultLogger) Error(format string, v ...any) {
	l.errorLogger.Printf(format, v...)
}

func (l *defaultLogger) Fatal(format string, v ...any) {
	l.fatalLogger.Fatalf(format, v...)
}

func (l *defaultLogger) Panic(format string, v ...any) {
	l.panicLogger.Panicf(format, v...)
}
