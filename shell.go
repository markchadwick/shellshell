package main

import (
	"io"
)

type Shell interface {
	Prompt() string
	Close() error
	CloseWithError(error) error
	Handle([]string, io.Reader, io.Writer, io.Writer) error
	Complete(string) []string
}
