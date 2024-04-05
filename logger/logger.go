package logger

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"
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
	Service  string   `json:"service"`   // Service name
	ProdMode bool     `json:"prod_mode"` // Production mode
}

// Logger is the main struct for creating and managing logs
type Logger struct {
	options    *LogOptions
	fileWriter *os.File
	template   *template.Template
}

func (l *Logger) init() error {
	// Define the template string with padding
	tmpl := `{{.Time}} | {{.Level}} | {{.Service}} | {{printf "%-20s" .Msg}}{{indent 4 (printf "%v" .Data)}}`

	// Define custom indent function
	funcMap := template.FuncMap{
		"indent": func(i int, input string) string {
			padding := strings.Repeat(" ", i)
			return padding + input
		},
	}

	// Parse the template with custom function map
	var err error
	l.template, err = template.New("logMessage").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return fmt.Errorf("error parsing log template: %w", err)
	}
	return nil
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
func (l *Logger) log(level LogLevel, msg string, data ...interface{}) {
	if level < l.options.Level {
		return
	}

	if l.options.ProdMode && level == Debug {
		return
	}

	var levelString string
	switch level {
	case Debug:
		levelString = "debug"
	case Info:
		levelString = "info"
	case Warning:
		levelString = "warning"
	case Error:
		levelString = "error"
	}

	// Calculate the padding based on the length of the log level
	padding := strings.Repeat(" ", 10-len(levelString))

	// Check if template is initialized before using it
	if l.template == nil {
		err := l.init()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing log template: %v\n", err)
			return
		}
	}

	strData := make([]string, 0, len(data)) // Create an empty string slice
	for _, element := range data {
		strData = append(strData, fmt.Sprint(element)) // Convert and append
	}

	// Join the string slice
	dataStr := strings.Join(strData, " ")

	// Use the template to format the message
	var buffer bytes.Buffer
	if err := l.template.Execute(&buffer, map[string]string{
		"Time":    time.Now().Format("2006-01-02 15:04:05"), // Current time in format
		"Level":   levelString + padding,                    // Add the padding to the log level
		"Service": l.options.Service,
		"Msg":     msg,
		"Data":    dataStr, // Pass the string
	}); err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting log message: %v\n", err)
		return
	}

	// Always append a newline for consistent formatting
	message := buffer.String() + "\n"

	// Log to configured destinations with explicit error handling
	if l.options.ToStdout {
		if _, err := fmt.Fprint(os.Stdout, message); err != nil {
			l.Errorf("Error writing log to stdout: %v", err)
		}
	}

	if l.options.ToFile {
		if _, err := l.fileWriter.Write([]byte(message)); err != nil {
			l.Errorf("Error writing log to file: %v", err)
		}
	}
}

// Debug logs a message with Debug level
func (l *Logger) Debug(msg string, data ...interface{}) {
	l.log(Debug, msg, data...)
}

func (l *Logger) Info(msg string, data ...interface{}) {
	l.log(Info, msg, data...)
}

func (l *Logger) Warning(msg string, data ...interface{}) {
	l.log(Warning, msg, data...)
}

func (l *Logger) Error(msg string, data ...interface{}) {
	l.log(Error, msg, data...)
}

// Close closes the log file if enabled
func (l *Logger) Close() error {
	if l.fileWriter == nil {
		return nil
	}
	if err := l.fileWriter.Close(); err != nil {
		return fmt.Errorf("error closing log file: %w", err)
	}
	l.fileWriter = nil
	return nil
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.log(Info, fmt.Sprintf(format, a...))
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.log(Error, fmt.Sprintf(format, a...))
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.log(Debug, fmt.Sprintf(format, a...))
}

func (l *Logger) Warningf(format string, a ...interface{}) {
	l.log(Warning, fmt.Sprintf(format, a...))
}

// Printf directly logs a formatted message with no level or data
func (l *Logger) Printf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...) + "\n"
	if l.options.ToStdout {
		if _, err := fmt.Fprint(os.Stdout, message); err != nil {
			l.Errorf("Error writing log to stdout: %v", err)
		}
	}

	if l.options.ToFile {
		if _, err := l.fileWriter.Write([]byte(message)); err != nil {
			l.Errorf("Error writing log to file: %v", err)
		}
	}
}
