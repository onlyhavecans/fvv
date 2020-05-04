# fvv - Fusozay Var Var

A CLI tool for quick text template rendering

Fusozay Var Var means "have fun"
It is a reference to something I see a lot

Fusozay Var Var is a CLI application for quickly rendering out text templates.
I often write outfit and character descriptions that reuses a lot of elements.
This allows me to DRY up my descriptions and still quickly get results.

Template requirements:
- All the templates must be valid golang templates.
- They must not require any variables to be passed in.
- All templates must have the '.tmpl' extension. All other files ignored.
- The template you want to render must be [named definitions](https://golang.org/pkg/text/template/#hdr-Nested_template_definitions).


## Usage

All errors go to stderr so you can pipe this to something like pbcopy

1. Set up a quantity of templates
1. run `fvv <template definition>` from directory containing template(s)
