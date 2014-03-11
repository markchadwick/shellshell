package shellshell

import (
	"fmt"
	"io"
	"strings"
)

type Handler func(string, io.Reader, io.Writer, io.Writer) error

type Builder struct {
	completer *Completer
	handlers  map[string]Handler
	prompt    func() string
	// prompt         func() string
	// close          func() error
	// closeWithError func(error) error
	// commands       map[string]Command
}

var _ Shell = new(Builder)

func New() *Builder {
	return &Builder{
		completer: NewCompleter(),
		handlers:  make(map[string]Handler),
	}
}

func (b *Builder) Prompt() string {
	if b.prompt != nil {
		return b.prompt()
	}
	return ""
}

func (b *Builder) Handle(l string, i io.Reader, o, e io.Writer) error {
	if h, ok := b.handlers[strings.Fields(l)[0]]; ok {
		return h(l, i, o, e)
	}
	return fmt.Errorf("No handler for %s", l)
}

func (b *Builder) Complete(l string) []string {
	return b.completer.Complete(l)
}

func (b *Builder) SetPrompt(f func() string) *Builder {
	b.prompt = f
	return b
}

func (b *Builder) Handler(cmd string, h Handler) *Builder {
	b.handlers[cmd] = h
	b.completer.Prefix(cmd, func(string) []string { return []string{cmd + " "} })
	return b
}

/*
func NewBuilder() *Builder {
	return &Builder{
		prompt:         func() string { return "prompt> " },
		close:          func() error { return io.EOF },
		closeWithError: func(e error) error { return e },
		commands:       make(map[string]Command),
	}
}



func (b *Builder) Prompt() string {
	return b.prompt()
}

func (b *Builder) Close() error {
	return b.close()
}

func (b *Builder) CloseWithError(e error) error {
	return b.closeWithError(e)
}

func (b *Builder) Handle(cmd string, i io.Reader, o io.Writer, e io.Writer) error {
	fields := strings.Fields(cmd)
	if len(fields) < 1 {
		return nil
	}

	if c, ok := b.commands[fields[0]]; ok {
		return c(cmd, i, o, e)
	}
	return nil
}

func (b *Builder) Complete(s string) []string {
	completions := make([]string, 0)
	for c, _ := range b.commands {
		if strings.HasPrefix(c, s) {
			completions = append(completions, c)
		}
	}
	return completions
}

// ----------------------------------------------------------------------------
// Builder functions

func (b *Builder) WithPrompt(f func() string) *Builder {
	b.prompt = f
	return b
}

func (b *Builder) Cmd(name string, c Command) *Builder {
	b.commands[name] = c
	return b
}
*/
