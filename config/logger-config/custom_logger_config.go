package logger_config

import (
	"fmt"
	"strings"
)

type CustomLogger struct{}

// Optional: Suppress this specific noisy log
func shouldLog(message string) bool {
	return !strings.Contains(message, "json call result data")
}

// Required method for internal logs (like RPC bridge messages)
func (l *CustomLogger) Print(message string) {
	if shouldLog(message) {
		fmt.Println(message)
	}
}

// Standard logger-config interface methods
func (l *CustomLogger) Trace(message string) {
	if shouldLog(message) {
		fmt.Println("[TRACE]", message)
	}
}

func (l *CustomLogger) Debug(message string) {
	if shouldLog(message) {
		fmt.Println("[DEBUG]", message)
	}
}

func (l *CustomLogger) Info(message string) {
	if shouldLog(message) {
		fmt.Println("[INFO]", message)
	}
}

func (l *CustomLogger) Warning(message string) {
	if shouldLog(message) {
		fmt.Println("[WARNING]", message)
	}
}

func (l *CustomLogger) Error(message string) {
	if shouldLog(message) {
		fmt.Println("[ERROR]", message)
	}
}

func (l *CustomLogger) Fatal(message string) {
	if shouldLog(message) {
		fmt.Println("[FATAL]", message)
	}
}
