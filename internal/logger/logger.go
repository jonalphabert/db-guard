package logger

import (
	"log"
	"os"
	"path/filepath"
)

var (
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	warnLogger    *log.Logger
	successLogger *log.Logger
)

func Init(logPath string) error {
	// Ensure the directory exists
	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(
		logPath,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}

	infoLogger = log.New(file, "[INFO] ", log.Ldate|log.Ltime)
	warnLogger = log.New(file, "[WARN] ", log.Ldate|log.Ltime)
	errorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime)
	successLogger = log.New(file, "[SUCCESS] ", log.Ldate|log.Ltime)

	return nil
}

func Info(format string, v ...interface{}) {
	if infoLogger != nil {
		infoLogger.Printf(format, v...)
	}
}

func Warn(format string, v ...interface{}) {
	if warnLogger != nil {
		warnLogger.Printf(format, v...)
	}
}

func Error(format string, v ...interface{}) {
	if errorLogger != nil {
		errorLogger.Printf(format, v...)
	}
}

func Success(format string, v ...interface{}) {
	if successLogger != nil {
		successLogger.Printf(format, v...)
	}
}
