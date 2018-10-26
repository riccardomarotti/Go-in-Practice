package main

import (
	"fmt"
)

type ParseError struct {
	Message    string
	Line, Char int
}

func (p *ParseError) Error() string {
	format := "%s on line %d, char %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}
