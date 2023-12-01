package utils

import (
	"fmt"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

var Instance *LogManager

type LogManager struct {
	logLevel LogLevel
}

func NewLogManager() *LogManager {
	Instance = &LogManager{
		logLevel: Debug,
	}
	return Instance
}

func (lm *LogManager) SetLogLevel(level LogLevel) {
	lm.logLevel = level
}

func (lm *LogManager) Debug(log string) {
	if lm.logLevel > Debug {
		return
	}
	fmt.Println("[Deb]", log)
}

func (lm *LogManager) Info(log string) {
	if lm.logLevel > Info {
		return
	}
	fmt.Println("[Inf]", log)
}

func (lm *LogManager) Warn(log string) {
	if lm.logLevel > Warn {
		return
	}
	fmt.Println("[Wrn]", log)
}

func (lm *LogManager) Error(log string) {
	if lm.logLevel > Error {
		return
	}
	fmt.Println("[Err]", log)
}

func (lm *LogManager) Fatal(log string) {
	if lm.logLevel > Fatal {
		return
	}
	fmt.Println("[Fat]", log)
}
