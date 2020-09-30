package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"text/template"
)

const (
	templateGlob = "*.tmpl"
	exitFail     = 1
	usage        = `Quick Usage: fvv <template definition name>

Fusozay Var Var is a CLI application for quickly rendering out text templates.
I often write outfit and character descriptions that reuses a lot of elements.
This allows me to DRY up my descriptions and still quickly get results.

Template requirements:
- All of the templates must be valid golang templates.
- They must not require any variables to be passed in.
- All templates must have the '.tmpl' extension. All other files ignored.
- The template you want to render must be "named"
https://golang.org/pkg/text/template/#hdr-Nested_template_definitions

Version 1.0.0
`
)

func main() {
	if err := run(os.Args, templateGlob, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, fileGlob string, stdout io.Writer) error {
	if len(args) < 2 {
		return errors.New(usage)
	}

	tmpl, err := template.ParseGlob(fileGlob)
	if err != nil {
		return fmt.Errorf("checking for templates: %v", err)
	}

	err = tmpl.ExecuteTemplate(stdout, args[1], "no data needed")
	if err != nil {
		return fmt.Errorf("executing template %s: %v", args[1], err)
	}

	return nil
}
