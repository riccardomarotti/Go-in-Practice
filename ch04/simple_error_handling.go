package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if result, err := concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result)
	}
}

func concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No string supplied")
	}

	return strings.Join(parts, " "), nil
}
