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
	"strings"
	"fmt"
)

type level uint32

const (
	fatalLevel level = iota
	errorLevel
	infoLevel
	debugLevel
)

var levels = []level{fatalLevel, errorLevel, infoLevel, debugLevel}

func (lvl level) String() string {
	switch lvl {
	case fatalLevel:
		return "FATAL"
	case errorLevel:
		return "ERROR"
	case infoLevel:
		return "INFO"
	case debugLevel:
		return "DEBUG"
	}

	return "unknown"
}

func parseLevel(s string) (level, error) {
	switch strings.ToLower(s) {
	case "fatal":
		return fatalLevel, nil
	case "error":
		return errorLevel, nil
	case "info":
		return infoLevel, nil
	case "debug":
		return debugLevel, nil
	}

	var lvl level
	return lvl, fmt.Errorf("not a valid level: %q", s)
}
