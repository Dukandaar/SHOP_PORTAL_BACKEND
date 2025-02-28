package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	logLevel := strings.ToUpper(entry.Level.String())
	message := entry.Message

	// ANSI escape codes for color (only used if DisableColors is false)
	infoColor := "\x1b[34m"
	timestampColor := "\x1b[32m"
	prefixColor := "\x1b[35m"
	resetColor := "\x1b[0m"

	// Split the message to separate the prefix and the rest
	parts := strings.SplitN(message, ":", 2)
	prefix := parts[0] + ":"
	rest := ""
	if len(parts) > 1 {
		rest = parts[1]
	}

	var formatted string
	var DisableColors = os.Getenv("DISABLE_COLORS")

	if DisableColors == "false" {
		// Colors for local development
		formatted = fmt.Sprintf("%s%s%s[%s%s%s] %s%s%s%s", infoColor, logLevel, resetColor, timestampColor, timestamp, resetColor, prefixColor, prefix, resetColor, rest)
	} else {
		// No colors for CloudWatch
		formatted = fmt.Sprintf("%s[%s] %s%s", logLevel, timestamp, prefix, rest)
	}

	b.WriteString(formatted + "\n")

	return b.Bytes(), nil
}

var Logger *logrus.Logger

func NewLogger() {
	if Logger != nil {
		return
	}

	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)

	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(new(CustomFormatter))
}
