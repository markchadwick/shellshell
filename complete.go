package shellshell

import (
	"strings"
)

type Completion interface {
	Matches(string) bool
	Complete(l string) []string
}

type Completer struct {
	completions []Completion
}

func NewCompleter() *Completer {
	return &Completer{}
}

func (c *Completer) Complete(line string) []string {
	results := make([]string, 0)

	for _, completion := range c.completions {
		if completion.Matches(line) {
			results = append(results, completion.Complete(line)...)
		}
	}

	return results
}

func (c *Completer) Add(completion Completion) {
	c.completions = append(c.completions, completion)
}

// ----------------------------------------------------------------------------
// Completions

type prefixCompletion struct {
	prefix   string
	complete func(string) []string
}

var _ Completion = new(prefixCompletion)

func (pc *prefixCompletion) Matches(l string) bool {
	return strings.HasPrefix(pc.prefix, l)
}

func (pc *prefixCompletion) Complete(l string) []string {
	return pc.complete(l)
}

func (c *Completer) Prefix(p string, f func(string) []string) {
	c.Add(&prefixCompletion{
		prefix:   p,
		complete: f,
	})
}
