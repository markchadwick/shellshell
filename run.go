package shellshell

/*
import (
	"os"

	"github.com/GeertJohan/go.linenoise"
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
	err = r.shell.Handle(line, os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		linenoise.AddHistory(line)
	}
	return err
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
*/
