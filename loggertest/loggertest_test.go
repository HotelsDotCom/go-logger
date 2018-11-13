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

package loggertest

import (
	"github.com/stretchr/testify/assert"
	"github.com/HotelsDotCom/go-logger"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPanicWhenFatalIsLogged(t *testing.T) {

	Init(LogLevelError)
	defer Reset()

	assert.Panics(t, func() { logger.Fatal("Panic!") })
}

func TestClearLogMessages(t *testing.T) {

	Init(LogLevelInfo)
	defer Reset()

	logger.Error("messages before clear")
	ClearLogMessages()
	logger.Error("messages after clear")

	logMessages := GetLogMessages()
	assert.Equal(t, 1, len(logMessages))
	assert.Contains(t, logMessages[0].RawMessage, "messages after clear")
}

func TestLogMessageStructIsPopulateFromLog(t *testing.T) {

	Init(LogLevelDebug)
	defer Reset()

	logger.Debug("test debug message")

	logMessges := GetLogMessages()
	require.Equal(t, 1, len(logMessges))
	assert.Contains(t, logMessges[0].RawMessage, "[DEBUG]")
	assert.Contains(t, logMessges[0].RawMessage, "test debug message")
	assert.Equal(t, "DEBUG", logMessges[0].Level)
	assert.Equal(t, "test debug message", logMessges[0].Message)
}


func TestDoesntPanicWhenThereareNoMessages(t *testing.T) {

	Init(LogLevelDebug)
	defer Reset()

	assert.Empty(t, GetLogMessages())
}