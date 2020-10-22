package logger

import (
	"fmt"

	"github.com/fatih/color"
)

const format = "%s  %s\n"

// Error ...
func Error(err error) {
	badge := color.New(color.FgRed).SprintFunc()
	fmt.Printf(format, badge("ERROR"), err.Error())
}

// Info ...
func Info(msg string) {
	badge := color.New(color.FgCyan).SprintFunc()
	fmt.Printf(format, badge("INFO"), msg)
}

// Warn ...
func Warn(msg string) {
	badge := color.New(color.FgYellow).SprintFunc()
	fmt.Printf(format, badge("WARN"), msg)
}

// Debug ...
func Debug(msg string) {
	badge := color.New(color.FgHiBlack).SprintFunc()
	fmt.Printf(format, badge("DEBUG"), msg)
}

// Success ...
func Success(msg string) {
	badge := color.New(color.FgGreen).SprintFunc()
	fmt.Printf(format, badge("SUCCESS"), msg)
}
