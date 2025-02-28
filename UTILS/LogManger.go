package utils

import (
	structs "SHOP_PORTAL_BACKEND/STRUCTS"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/kataras/iris/v12"
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

// LogRequest logs all relevant request information
func LogRequest(logPrefix string, ctx iris.Context, reqBody interface{}) {
	headers := ctx.Request().Header
	url := ctx.Request().URL.String()
	method := ctx.Request().Method
	remoteAddr := ctx.RemoteAddr()
	queryParams := ctx.Request().URL.Query()

	logData := structs.RequestLogData{
		Headers:     headers,
		QueryParams: queryParams,
		RequestBody: reqBody,
		Method:      method,
		URL:         url,
		RemoteAddr:  remoteAddr,
	}

	logJSON, err := json.Marshal(logData)
	if err != nil {
		Logger.WithError(err).Error(logPrefix + "Failed to marshal request log data")
		return
	}

	Logger.Infof(logPrefix+"Request Details: %s", logJSON)
}

func LogResponse(logPrefix string, response interface{}) {
	logJSON, err := json.Marshal(response)
	if err != nil {
		Logger.WithError(err).Error(logPrefix + "Failed to marshal response to JSON for logging")
		return
	}

	Logger.Infof(logPrefix+"Response: %s", logJSON)
}
