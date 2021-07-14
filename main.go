package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"text/template"
)

const (
	exitFail    = 1
	usageFooter = `
Fusozay Var Var is a CLI pre-processor for rendering out golang txt templates.
I often write outfit and character descriptions that reuses a lot of elements.
This allows me to DRY up my descriptions and still quickly get results.

Template requirements:
- All of the templates must be valid golang templates.
- They must not require any variables to be passed in.
- The template you want to render must be "named"
https://golang.org/pkg/text/template/#hdr-Nested_template_definitions

Version 2.0.0
`
)

func main() {
	if err := verifyCharDevice(os.Stdin); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(exitFail)
	}

	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(exitFail)
	}
}

func verifyCharDevice(stdin *os.File) error {
	fi, err := stdin.Stat() // get the FileInfo struct describing the standard input.
	if err != nil {
		return fmt.Errorf("doing stat on %p: %v", stdin, err)
	}
	if (fi.Mode() & os.ModeCharDevice) != 0 {
		return errors.New("stdin was not passed to the app")
	}
	return nil
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprintf(flags.Output(), "Quick Usage: %s < <template>\n", args[0])
		flags.PrintDefaults()
		fmt.Fprint(flags.Output(), usageFooter)
	}
	var (
	// glob = flags.String("t", fileGlob, "glob to use to find templates")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if flags.NArg() != 0 {
		flags.Usage()
		return errors.New("pass template in stdin")
	}

	f, err := io.ReadAll(stdin)
	if err != nil {
		return fmt.Errorf("reading from stdin: %v", err)
	}

	tmpl, err := template.New("name").Parse(string(f))
	if err != nil {
		return fmt.Errorf("parsing template: %v", err)
	}

	err = tmpl.Execute(stdout, "no data needed")
	if err != nil {
		return fmt.Errorf("executing template %s: %v", args[1], err)
	}

	return nil
}
