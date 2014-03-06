package main

import (
	"io"
	"log"
)

type builder struct {
	prompt         func() string
	close          func() error
	closeWithError func(error) error
}

func New() *builder {
	return &builder{
		prompt:         func() string { return "prompt> " },
		close:          func() error { return io.EOF },
		closeWithError: func(e error) error { return e },
	}
}

var _ Shell = new(builder)

// ----------------------------------------------------------------------------
// Shell interface

func (b *builder) Prompt() string {
	return b.prompt()
}

func (b *builder) Close() error {
	return b.close()
}

func (b *builder) CloseWithError(e error) error {
	return b.closeWithError(e)
}

func (b *builder) Handle(l []string, i io.Reader, o io.Writer, e io.Writer) error {
	return nil
}

func (b *builder) Complete(s string) []string {
	log.Printf("Completing: '%s'", s)
	return nil
}

// ----------------------------------------------------------------------------
// Builder functions

func (b *builder) WithPrompt(f func() string) *builder {
	b.prompt = f
	return b
}
