package utils

import "fmt"

type Logger struct{}

func (l *Logger) Info(message string) {
	fmt.Printf("[INFO]: %s", message)
}
