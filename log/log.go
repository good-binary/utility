package customlog

import (
	"encoding/json"
	"fmt"
	"os"
)

// LogLevel defines the severity level of a log message
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

// LogOptions holds configuration options for the logger
type LogOptions struct {
	ToFile   bool     `json:"to_file"`   // Enable logging to file
	ToStdout bool     `json:"to_stdout"` // Enable logging to terminal
	JSON     bool     `json:"json"`      // Enable JSON formatting
	Level    LogLevel `json:"level"`     // Minimum severity level to log
	LogFile  string   `json:"log_file"`  // Path to the log file (optional)
}

// Logger is the main struct for creating and managing logs
type Logger struct {
	options    *LogOptions
	fileWriter *os.File
}

// NewLogger creates a new Logger instance with the provided options
func NewLogger(opts *LogOptions) (*Logger, error) {
	if opts == nil {
		opts = &LogOptions{
			ToStdout: true,
			Level:    Info,
		}
	}

	logger := &Logger{
		options: opts,
	}

	if opts.ToFile {
		var err error
		logger.fileWriter, err = os.OpenFile(opts.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("error opening log file: %w", err)
		}
	}

	return logger, nil
}

// log writes a message to the configured destinations (file, terminal, or both)
func (l *Logger) log(level LogLevel, msg string, data interface{}) {
	if level < l.options.Level {
		return
	}

	prefix := fmt.Sprintf("[%v] ", level)
	message := fmt.Sprintf("%s%s\n", prefix, msg)

	if l.options.JSON {
		if data != nil {
			jsonData, err := json.Marshal(data)
			if err != nil {
				message += fmt.Sprintf("Error marshalling data: %v\n", err)
			} else {
				message += string(jsonData) + "\n"
			}
		}
	} else {
		if data != nil {
			message += fmt.Sprintf("%v\n", data)
		}
	}

	if l.options.ToStdout {
		fmt.Print(message)
	}

	if l.options.ToFile {
		l.fileWriter.Write([]byte(message))
	}
}

// Debug logs a message with Debug level
func (l *Logger) Debug(msg string, data ...interface{}) {
	l.log(Debug, msg, data)
}

// Info logs a message with Info level
func (l *Logger) Info(msg string, data ...interface{}) {
	l.log(Info, msg, data)
}

// Warning logs a message with Warning level
func (l *Logger) Warning(msg string, data ...interface{}) {
	l.log(Warning, msg, data)
}

// Error logs a message with Error level
func (l *Logger) Error(msg string, data ...interface{}) {
	l.log(Error, msg, data)
}

// Close closes the log file if enabled
func (l *Logger) Close() error {
	if l.fileWriter != nil {
		return l.fileWriter.Close()
	}
	return nil
}
