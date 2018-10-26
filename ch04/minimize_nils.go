package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	result, _ := concat(args...)
	fmt.Printf("Concatenated string: '%s'\n", result)
}

func concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No string supplied")
	}

	return strings.Join(parts, " "), nil
}
