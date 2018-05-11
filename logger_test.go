/*
Copyright (C) 2016-2018 Expedia Inc.

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

package logger

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testLogWriter = &bytes.Buffer{}

func beforeTest(logLevel string) {
	SetLevel(logLevel)
	SetOutput(testLogWriter)
	SetExitFn(func() {})
}

func afterTest() {
	testLogWriter.Reset()
}

func TestLogger_DefaultLogLevel(t *testing.T) {

	assert.Equal(t, "INFO", std.logLevel.String())
}

func TestLogger_SettingDEBUGLogLevelEnablesDEBUG_INFO_ERROR_FATAL(t *testing.T) {

	beforeTest("DEBUG")
	defer afterTest()

	Debug("debug message")
	Info("info message")
	Error("error message")
	Fatal("fatal message")

	logMessages := testLogWriter.String()
	assert.Contains(t, logMessages, "debug message")
	assert.Contains(t, logMessages, "info message")
	assert.Contains(t, logMessages, "error message")
	assert.Contains(t, logMessages, "fatal message")
}

func TestLogger_SettingINFOLogLevelEnablesINFO_ERROR_FATAL(t *testing.T) {

	beforeTest("INFO")
	defer afterTest()

	Debug("debug message")
	Info("info message")
	Error("error message")
	Fatal("fatal message")

	logMessages := testLogWriter.String()
	assert.NotContains(t, logMessages, "debug message")
	assert.Contains(t, logMessages, "info message")
	assert.Contains(t, logMessages, "error message")
	assert.Contains(t, logMessages, "fatal message")
}

func TestLogger_SettingERRORLogLevelEnablesERROR_FATAL(t *testing.T) {

	beforeTest("ERROR")
	defer afterTest()

	Debug("debug message")
	Info("info message")
	Error("error message")
	Fatal("fatal message")

	logMessages := testLogWriter.String()
	assert.NotContains(t, logMessages, "debug message")
	assert.NotContains(t, logMessages, "info message")
	assert.Contains(t, logMessages, "error message")
	assert.Contains(t, logMessages, "fatal message")
}

func TestLogger_SettingFATALogLevelEnablesOnly_FATAL(t *testing.T) {

	beforeTest("FATAL")
	defer afterTest()

	Debug("debug message")
	Info("info message")
	Error("error message")
	Fatal("fatal message")

	logMessages := testLogWriter.String()
	assert.NotContains(t, logMessages, "debug message")
	assert.NotContains(t, logMessages, "info message")
	assert.NotContains(t, logMessages, "error message")
	assert.Contains(t, logMessages, "fatal message")
}

func TestLoggerMessageIsInCorrectFormat(t *testing.T) {

	beforeTest("DEBUG")
	defer afterTest()

	Debug("testing debug log message")

	messageFormat := `\[%s\] [\d]{4}\/[\d]{2}\/[\d]{2} [\d]{2}:[\d]{2}:[\d]{2} main\.go:[\d]{1,}: %s`
	expectedMessage := fmt.Sprintf(messageFormat, "DEBUG", "testing debug log message")
	assert.Regexp(t, expectedMessage, testLogWriter.String())
}
