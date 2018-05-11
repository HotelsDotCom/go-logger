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
	"fmt"
	"io"
	"log"
	"sync"
)

type logger struct {
	mu        sync.Mutex
	out       io.Writer
	logLevel  level
	calldepth int
	loggers   map[level]*log.Logger
	exitFn    func()
}

func (l logger) log(lvl level, v ...interface{}) {
	if l.logLevel >= lvl {
		l.output(lvl, fmt.Sprint(v...))
	}
}

func (l logger) logf(lvl level, format string, v ...interface{}) {
	if l.logLevel >= lvl {
		l.output(lvl, fmt.Sprintf(format, v...))
	}
}

func (l logger) output(lvl level, s string) {
	l.loggers[lvl].Output(l.calldepth, s)
}

func createLoggers(out io.Writer) map[level]*log.Logger {
	loggers := map[level]*log.Logger{}

	for _, lvl := range levels {
		prefix := fmt.Sprintf("[%s] ", lvl.String())
		loggers[lvl] = log.New(out, prefix, log.Ldate|log.Ltime|log.Lshortfile)
	}

	return loggers
}
