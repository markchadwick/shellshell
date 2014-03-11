package shellshell

import (
	"io"
	"os"

	"github.com/GeertJohan/go.linenoise"
)

type Shell interface {
	Prompt() string
	// Close() error
	// CloseWithError(error) error
	Handle(string, io.Reader, io.Writer, io.Writer) error
	Complete(string) []string
}

func Run(s Shell) error {
	i := os.Stdin
	o := os.Stdout
	e := os.Stderr

	linenoise.SetCompletionHandler(s.Complete)
	for {
		if err := loop(s, i, o, e); err != nil {
			return err
		}
	}
	return nil
}

func loop(s Shell, i io.Reader, o, e io.Writer) error {
	line, err := linenoise.Line(s.Prompt())
	if err != nil {
		return err
	}
	if err = s.Handle(line, i, o, e); err != nil {
		return err
	}
	linenoise.AddHistory(line)
	return nil
}
