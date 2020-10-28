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
	templateGlob = "*.tmpl"
	exitFail     = 1
	usageFooter  = `
Fusozay Var Var is a CLI application for quickly rendering out text templates.
I often write outfit and character descriptions that reuses a lot of elements.
This allows me to DRY up my descriptions and still quickly get results.

Template requirements:
- All of the templates must be valid golang templates.
- They must not require any variables to be passed in.
- The template you want to render must be "named"
https://golang.org/pkg/text/template/#hdr-Nested_template_definitions

Version 1.2.0
`
)

func main() {
	if err := run(os.Args, templateGlob, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, fileGlob string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() {
		fmt.Fprintf(flags.Output(), "Quick Usage: %s <template definition name>\n", args[0])
		flags.PrintDefaults()
		fmt.Fprint(flags.Output(), usageFooter)
	}
	var (
		glob = flags.String("t", fileGlob, "glob to use to find templates")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if flags.NArg() != 1 {
		flags.Usage()
		return errors.New("template definition missing")
	}
	nonFlagArg := flags.Args()[0]


	tmpl, err := template.ParseGlob(*glob)
	if err != nil {
		return fmt.Errorf("checking for templates: %v", err)
	}

	err = tmpl.ExecuteTemplate(stdout, nonFlagArg, "no data needed")
	if err != nil {
		return fmt.Errorf("executing template %s: %v", args[1], err)
	}

	return nil
}
