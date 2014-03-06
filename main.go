package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

type EchoShell struct {
	i int
}

func (es *EchoShell) Prompt() string {
	es.i += 1
	return fmt.Sprintf("echo %d> ", es.i)
}

func (es *EchoShell) Echo(args []string, i io.Reader, o io.Writer, e io.Writer) error {
	res := strings.Join(args, " ")
	_, err := o.Write([]byte(res))
	return err
}

func main() {
	echo := new(EchoShell)
	builder := New().
		WithPrompt(echo.Prompt).
		WithCmd("echo", echo.Echo)

	if err := Run(builder); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}
	log.Printf("clean exit!")
}
