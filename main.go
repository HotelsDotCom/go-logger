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

package logger

import (
	"io"
	"os"
)

var std = newLogger()

func newLogger() *logger {
	out := os.Stderr
	return &logger{
		out:       out,
		logLevel:  infoLevel,
		calldepth: 3,
		loggers:   createLoggers(out),
		exitFn:    func() { os.Exit(1) },
	}
}

func SetOutput(out io.Writer) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.out = out
	for _, l := range std.loggers {
		l.SetOutput(out)
	}
}

func SetLevel(level string) {
	l, err := parseLevel(level)
	if err != nil {
		Error(err)
	}

	std.mu.Lock()
	defer std.mu.Unlock()
	std.logLevel = l
}

func SetExitFn(exitFn func()) {
	std.mu.Lock()
	defer std.mu.Unlock()
	std.exitFn = exitFn
}

func GetLogLevel() string {
	return std.logLevel.String()
}

func GetOut() io.Writer {
	return std.out
}

func GetExitFn() func() {
	return std.exitFn
}

func Debug(v ...interface{}) {
	std.log(debugLevel, v...)
}

func Debugf(format string, v ...interface{}) {
	std.logf(debugLevel, format, v...)
}

func Info(v ...interface{}) {
	std.log(infoLevel, v...)
}

func Infof(format string, v ...interface{}) {
	std.logf(infoLevel, format, v...)
}

func Error(v ...interface{}) {
	std.log(errorLevel, v...)
}

func Errorf(format string, v ...interface{}) {
	std.logf(errorLevel, format, v...)
}

func Fatal(v ...interface{}) {
	std.log(fatalLevel, v...)
	std.exitFn()
}

func Fatalf(format string, v ...interface{}) {
	std.logf(fatalLevel, format, v...)
	std.exitFn()
}
