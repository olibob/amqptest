package utilities

import (
	"fmt"
	"log"
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
