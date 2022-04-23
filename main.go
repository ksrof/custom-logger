package main

import (
	"errors"
	"fmt"
	"strings"
)

// logger represents the data fields of the logger.
type logger struct {
	fields []string
}

// withFields appends to a slice the given fields as key-value pairs.
func (log *logger) withFields(fields map[string]string) *logger {
	for key, value := range fields {
		log.fields = append(log.fields, fmt.Sprintf("%s=%s", key, value))
	}

	return log
}

// print displays the logger information with a given kind and message.
func (log *logger) print(kind, msg string) {
	switch strings.ToLower(kind) {
	case "info":
		fmt.Printf("[%s] %s %s\n", strings.ToUpper(kind), msg, strings.Join(log.fields, ", "))
	case "warn":
		fmt.Printf("[%s] %s %s\n", strings.ToUpper(kind), msg, strings.Join(log.fields, ", "))
	case "erro":
		fmt.Printf("[%s] %s %s\n", strings.ToUpper(kind), msg, strings.Join(log.fields, ", "))
	case "fatl":
		fmt.Printf("[%s] %s %s\n", strings.ToUpper(kind), msg, strings.Join(log.fields, ", "))
	}

	log.fields = nil
}

func main() {
	// initialize the logger
	logger := &logger{}

	// create an error on purpose
	err := errors.New("keyboard not responding, press any key to continue")
	if err != nil {
		logger.withFields(map[string]string{
			"file":  "main.go",
			"line":  "46",
			"error": err.Error(),
		}).print("erro", "Something went wrong.")
	}
}
