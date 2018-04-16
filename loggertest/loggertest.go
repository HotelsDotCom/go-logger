/*
Copyright (C) 2018 Expedia Group.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// loggertest package provides test logger that provides function to inspect log messages and
// changes the default behaviour of fatal log (os.Exit) to panic. Tests are then able to recover
package loggertest

import (
	"bytes"
	"io"
	"github.com/HotelsDotCom/go-logger"
	"strings"
)

const (
	LogLevelDebug = "DEBUG"
	LogLevelInfo  = "INFO"
	LogLevelError = "ERROR"
	LogLevelFatal = "FATAL"
)

var (
	prevLogLevel   string
	prevWriter     io.Writer
	prevAfterFatal func(string)
	logWriter      *bytes.Buffer
)

// Overrides default logger with specified log level, swaps os.Exit() for panic() in case of fatal log.
// Logged messages can be inspected using GetLogMessages()
func Init(logLevel string) {

	prevLogLevel = logger.LogLevel
	prevWriter = logger.LogWriter
	prevAfterFatal = logger.AfterFatal
	logWriter = &bytes.Buffer{}

	logger.LogLevel = logLevel
	logger.LogWriter = logWriter
	logger.AfterFatal = func(msg string) { panic(msg) }

	logger.InitLoggers()
	ClearLogMessages()
}

type LogMessage struct {
	RawMessage string
	Message    string
	Level      string
}

func GetLogMessages() []LogMessage {

	messages := []LogMessage{}
	logs := strings.TrimSuffix(logWriter.String(), "\n")
	for _, line := range strings.Split(logs, "\n") {

		words := strings.SplitN(line, " ", 5)
		level := strings.Trim(words[0], "[]")

		logMessage := LogMessage{RawMessage: line, Message: words[4], Level: level}
		messages = append(messages, logMessage)
	}
	return messages
}

func ClearLogMessages() {
	logWriter.Reset()
}

func Reset() {

	ClearLogMessages()
	logger.LogLevel = prevLogLevel
	logger.LogWriter = prevWriter
	logger.AfterFatal = prevAfterFatal
	logger.InitLoggers()
}
