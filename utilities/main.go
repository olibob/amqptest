package utilities

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// FailOnError is san error wrapper
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// BodyFrom joins command line arguments in a string
func BodyFrom(args []string) string {
	s := ""
	if len(args) < 2 || args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

// BodyFrom2 joins command line arguments in a string
func BodyFrom2(args []string) string {
	s := ""
	if len(args) < 3 || args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

// SeverityFrom gets log Severity
func SeverityFrom(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}
