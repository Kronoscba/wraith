package logger

import (
	"fmt"
	"os"
	"time"
)

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
	DEBUG LogLevel = "DEBUG"
)

type Logger struct {
	file *os.File
}

func NewLogger(filename string) (*Logger, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &Logger{file: f}, nil
}

func (l *Logger) Log(level LogLevel, msg string) {
	timestamp := time.Now().Format(time.RFC3339)
	line := fmt.Sprintf("[%s] [%s] %s\n", timestamp, level, msg)
	
	fmt.Print(line) // stdout
	if l.file != nil {
		l.file.WriteString(line)
	}
}

func (l *Logger) Info(msg string)  { l.Log(INFO, msg) }
func (l *Logger) Warn(msg string)  { l.Log(WARN, msg) }
func (l *Logger) Error(msg string) { l.Log(ERROR, msg) }
func (l *Logger) Debug(msg string) { l.Log(DEBUG, msg) }

func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}
