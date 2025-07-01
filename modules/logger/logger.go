package logger

import (
	"log"
	"os"
	"sync"
)

type Logger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
	logOnce     sync.Once
}

// NewLogger creates a new Logger instance with initialized loggers.
func NewLogger() *Logger {
	return &Logger{}
}

// Init initializes the loggers for different levels.
func (l *Logger) Init() {
	l.logOnce.Do(func() {
		l.infoLogger = log.New(os.Stdout, "\033[34mINFO:\033[0m ", log.Ldate|log.Ltime|log.Lshortfile)
		l.warnLogger = log.New(os.Stdout, "\033[33mWARN:\033[0m ", log.Ldate|log.Ltime|log.Lshortfile)
		l.errorLogger = log.New(os.Stderr, "\033[31mERROR:\033[0m ", log.Ldate|log.Ltime|log.Lshortfile)
		l.fatalLogger = log.New(os.Stderr, "\033[35mFATAL:\033[0m ", log.Ldate|log.Ltime|log.Lshortfile)
	})
}

// Info logs an informational message.
func (l *Logger) Info(v ...interface{}) {
	l.Init() // Ensure loggers are initialized
	l.infoLogger.Println(v...)
}

// Warn logs a warning message.
func (l *Logger) Warn(v ...interface{}) {
	l.Init() // Ensure loggers are initialized
	l.warnLogger.Println(v...)
}

// Error logs an error message.
func (l *Logger) Error(v ...interface{}) {
	l.Init() // Ensure loggers are initialized
	l.errorLogger.Println(v...)
}

// Fatal logs a fatal message and exits the program.
func (l *Logger) Fatal(v ...interface{}) {
	l.Init() // Ensure loggers are initialized
	l.fatalLogger.Fatalln(v...)
}
