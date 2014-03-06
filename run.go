package main

import (
	"github.com/GeertJohan/go.linenoise"

	"log"
)

type runner struct {
	shell Shell
}

func (r *runner) Loop() error {
	line, err := linenoise.Line(r.shell.Prompt())
	if err != nil {
		if err == linenoise.KillSignalError {
			return r.shell.Close()
		}
		return r.shell.CloseWithError(err)
	}
	log.Printf("line: '%s'", line)
	return nil
}

func Run(s Shell) (err error) {
	runner := &runner{s}
	linenoise.SetCompletionHandler(runner.shell.Complete)
	for {
		if err := runner.Loop(); err != nil {
			return err
		}
	}
}
