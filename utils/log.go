package utils

import (
	"fmt"
)

const (
	//
	reset = "\033[0m"
	bold  = "\033[1m"

	// color
	red    = "\033[31m"
	yellow = "\033[33m"
	gray   = "\033[37m"
)

func Info(msg string, a ...any) {
	log(gray, "INFO", fmt.Sprintf(msg, a...))
}

func Warn(warning string, a ...any) {
	log(yellow, "WARN", fmt.Sprintf(warning, a...))
}

func Err(err error, a ...any) {
	log(red, "ERR ", fmt.Sprintf(err.Error(), a...))
}

func log(color, level string, message any) {
	fmt.Printf("[%s%s%s%s] %s\n", color, bold, level, reset, message)
}
