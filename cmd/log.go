package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// Logger struct
type Logger struct {
}

// NewLogger creates a new logger!
func NewLogger() *Logger {
	return &Logger{}
}

// Yellow - Outputs yellow text
func (l *Logger) Yellow(format string, a ...interface{}) {
	color.New(color.FgHiYellow).PrintfFunc()(format+"\n", a...)
}

// Yellowf - Outputs yellow text without the newline
func (l *Logger) Yellowf(format string, a ...interface{}) {
	color.New(color.FgHiYellow).PrintfFunc()(format, a...)
}

// Green - Outputs Green text
func (l *Logger) Green(format string, a ...interface{}) {
	color.New(color.FgHiGreen).PrintfFunc()(format+"\n", a...)
}

// White - Outputs White text
func (l *Logger) White(format string, a ...interface{}) {
	color.New(color.FgHiWhite).PrintfFunc()(format+"\n", a...)
}

// WhiteUnderline - Outputs White text with underline
func (l *Logger) WhiteUnderline(format string, a ...interface{}) {
	l.White(format, a...)
	l.White(l.underline(format))
}

// YellowUnderline - Outputs Yellow text with underline
func (l *Logger) YellowUnderline(format string, a ...interface{}) {
	l.Yellow(format, a...)
	l.Yellow(l.underline(format))
}

// underline returns a string of a line, the length of the message given to it
func (l *Logger) underline(message string) string {
	return strings.Repeat("-", len(message))
}

// Red - Outputs Red text
func (l *Logger) Red(format string, a ...interface{}) {
	color.New(color.FgHiRed).PrintfFunc()(format+"\n", a...)
}

// Error - Outputs an Error message
func (l *Logger) Error(format string, a ...interface{}) {
	color.New(color.FgHiRed).PrintfFunc()("Error: "+format+"\n", a...)
}

// PrintBanner prints the Wails banner before running commands
func (l *Logger) PrintBanner() error {
	banner1 := ` _       __      _ __    
| |     / /___ _(_) /____
| | /| / / __ ` + "`" + `/ / / ___/
| |/ |/ / /_/ / / (__  )  `
	banner2 := `|__/|__/\__,_/_/_/____/   `

	l.Yellowf(banner1)
	l.Red(Version)
	l.Yellowf(banner2)
	l.Green("https://wails.app")
	l.White("The lightweight framework for web-like apps")
	fmt.Println()

	return nil
}