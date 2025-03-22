package utils

import (
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

type Logger struct {
	loggers map[LogLevel]*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		loggers: map[LogLevel]*log.Logger{
			DEBUG: log.New(os.Stdout, colorBlue+"[DEBUG] "+colorReset, log.LstdFlags),
			INFO:  log.New(os.Stdout, colorGreen+"[INFO]  "+colorReset, log.LstdFlags),
			WARN:  log.New(os.Stdout, colorYellow+"[WARN]  "+colorReset, log.LstdFlags),
			ERROR: log.New(os.Stdout, colorRed+"[ERROR] "+colorReset, log.LstdFlags),
			FATAL: log.New(os.Stdout, colorRed+"[FATAL] "+colorReset, log.LstdFlags),
		},
	}
}

func (l *Logger) Debug(format string, v ...interface{}) {
	if l == nil || l.loggers[DEBUG] == nil {
		log.Println("[WARN] Logger is not initialized")
		return
	}
	l.loggers[DEBUG].Printf(format, v...)
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l == nil || l.loggers[INFO] == nil {
		log.Println("[WARN] Logger is not initialized")
		return
	}
	l.loggers[INFO].Printf(format, v...)
}

func (l *Logger) Warn(format string, v ...interface{}) {
	if l == nil || l.loggers[WARN] == nil {
		log.Println("[WARN] Logger is not initialized")
		return
	}
	l.loggers[WARN].Printf(format, v...)
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l == nil || l.loggers[ERROR] == nil {
		log.Println("[WARN] Logger is not initialized")
		return
	}
	l.loggers[ERROR].Printf(format, v...)
}

func (l *Logger) Fatal(format string, v ...interface{}) {
	if l == nil || l.loggers[FATAL] == nil {
		log.Println("[WARN] Logger is not initialized")
		os.Exit(1)
	}
	l.loggers[FATAL].Printf(format, v...)
	os.Exit(1)
}
