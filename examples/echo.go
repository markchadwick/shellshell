package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	shellshell "../"
)

func main() {
	var pwd = "/"

	shell := shellshell.
		New().
		SetPrompt(func() string {
		return fmt.Sprintf("%s> ", pwd)
	}).Handler("cd", func(l string, i io.Reader, o, e io.Writer) error {
		fs := strings.Fields(l)
		if len(fs) != 2 {
			fmt.Fprint(e, "Usage: cd <dir>\n")
			return nil
		}
		pwd = fs[1]
		return nil
	})

	if err := shellshell.Run(shell); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}

	log.Printf("clean exit!")
}
